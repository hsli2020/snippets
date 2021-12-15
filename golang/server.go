package pcopy

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"golang.org/x/time/rate"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

const (
	managerTickerInterval      = 30 * time.Second
	defaultMaxAuthAge          = time.Minute
	noAuthRequestAge           = 0
	rateLimitRequestsPerSecond = 2
	rateLimitBurstPerSecond    = 5
	rateLimitExpungeAfter      = 3 * time.Minute
	certCommonName             = "pcopy"
)

var (
	authOverrideParam   = "a"
	authHmacFormat      = "HMAC %d %d %s" // timestamp ttl b64-hmac
	authHmacRegex       = regexp.MustCompile(`^HMAC (\d+) (\d+) (.+)$`)
	authBasicRegex      = regexp.MustCompile(`^Basic (\S+)$`)
	clipboardPathFormat = "/%s"
	reservedFiles       = []string{"info", "verify", "join", "static"}

	//go:embed "web/index.gohtml"
	webTemplateSource string
	webTemplate       = template.Must(template.New("index").Funcs(templateFnMap).Parse(webTemplateSource))

	//go:embed web/static
	webStaticFs embed.FS

	//go:embed "scripts/join.sh.tmpl"
	joinTemplateSource string
	joinTemplate       = template.Must(template.New("join").Funcs(templateFnMap).Parse(joinTemplateSource))
)

// server is the main HTTP server struct. It's the one with all the good stuff.
type server struct {
	config       *Config
	countLimiter *limiter
	sizeLimiter  *limiter
	rateLimiter  map[string]*visitor
	routes       []route
	sync.Mutex
}

// visitor represents an API user, and its associated rate.Limiter used for rate limiting
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// infoResponse is the response returned by the / endpoint
type infoResponse struct {
	ServerAddr string `json:"serverAddr"`
	Salt       string `json:"salt"`
}

// handlerFnWithErr extends the normal http.HandlerFunc to be able to easily return errors
type handlerFnWithErr func(http.ResponseWriter, *http.Request) error

// route represents a HTTP route (e.g. GET /info), a regex that matches it and its handler
type route struct {
	method  string
	regex   *regexp.Regexp
	handler handlerFnWithErr
}

func newRoute(method, pattern string, handler handlerFnWithErr) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

// routeCtx is a marker struct used to find fields in route matches
type routeCtx struct{}

// webTemplateConfig is a struct defining all the things required to render the web root
type webTemplateConfig struct {
	Key              *Key
	KeyDerivIter     int
	KeyLenBytes      int
	FileExpireAfter  time.Duration
	CurlPinnedPubKey string
}

// Serve starts a server and listens for incoming HTTPS requests. The server handles all management operations (info,
// verify, ...), as well as the actual clipboard functionality (GET/PUT/POST). It also starts a background process
// to prune old
func Serve(config *Config) error {
	server, err := newServer(config)
	if err != nil {
		return err
	}
	go server.serverManager()
	return server.listenAndServeTLS()
}

func newServer(config *Config) (*server, error) {
	if config.ListenAddr == "" {
		return nil, errListenAddrMissing
	}
	if config.KeyFile == "" {
		return nil, errKeyFileMissing
	}
	if config.CertFile == "" {
		return nil, errCertFileMissing
	}
	if err := os.MkdirAll(config.ClipboardDir, 0700); err != nil {
		return nil, errClipboardDirNotWritable
	}
	if unix.Access(config.ClipboardDir, unix.W_OK) != nil {
		return nil, errClipboardDirNotWritable
	}
	return &server{
		config:       config,
		sizeLimiter:  newLimiter(config.ClipboardSizeLimit),
		countLimiter: newLimiter(int64(config.ClipboardCountLimit)),
		rateLimiter:  make(map[string]*visitor),
		routes:       nil,
	}, nil
}

func (s *server) listenAndServeTLS() error {
	if s.config.Key == nil {
		log.Printf("Listening on %s (UNPROTECTED CLIPBOARD)\n", s.config.ListenAddr)
	} else {
		log.Printf("Listening on %s\n", s.config.ListenAddr)
	}

	http.HandleFunc("/", s.handle)
	return http.ListenAndServeTLS(s.config.ListenAddr, s.config.CertFile, s.config.KeyFile, nil)
}

func (s *server) routeList() []route {
	if s.routes != nil {
		return s.routes
	}
	s.Lock()
	defer s.Unlock()

	s.routes = []route{
		newRoute("GET", "/", s.onlyIfWebUI(s.handleWebRoot)),
		newRoute("GET", "/static/.+", s.onlyIfWebUI(s.handleWebStatic)),
		newRoute("GET", "/info", s.limit(s.handleInfo)),
		newRoute("GET", "/verify", s.limit(s.auth(s.handleVerify))),
		newRoute("GET", "/join", s.limit(s.auth(s.handleJoin))),
		newRoute("GET", "/(?i)([a-z0-9][-_.a-z0-9]{1,100})", s.limit(s.auth(s.handleClipboardGet))),
		newRoute("PUT", "/(?i)([a-z0-9][-_.a-z0-9]{1,100})", s.limit(s.auth(s.handleClipboardPut))),
		newRoute("POST", "/(?i)([a-z0-9][-_.a-z0-9]{1,100})", s.limit(s.auth(s.handleClipboardPut))),
	}
	return s.routes
}

func (s *server) handle(w http.ResponseWriter, r *http.Request) {
	for _, route := range s.routeList() {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 && r.Method == route.method {
			ctx := context.WithValue(r.Context(), routeCtx{}, matches[1:])
			if err := route.handler(w, r.WithContext(ctx)); err != nil {
				if e, ok := err.(*errHTTPNotOK); ok {
					s.fail(w, r, e.code, e)
				} else {
					s.fail(w, r, http.StatusInternalServerError, err)
				}
			}
			return
		}
	}
	if r.Method == http.MethodGet {
		s.fail(w, r, http.StatusNotFound, errNoMatchingRoute)
	} else {
		s.fail(w, r, http.StatusBadRequest, errNoMatchingRoute)
	}
}

func (s *server) handleInfo(w http.ResponseWriter, r *http.Request) error {
	log.Printf("%s - %s %s", r.RemoteAddr, r.Method, r.RequestURI)

	salt := ""
	if s.config.Key != nil {
		salt = base64.StdEncoding.EncodeToString(s.config.Key.Salt)
	}

	response := &infoResponse{
		ServerAddr: s.config.ServerAddr,
		Salt:       salt,
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func (s *server) handleVerify(w http.ResponseWriter, r *http.Request) error {
	log.Printf("%s - %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	return nil
}

func (s *server) handleWebRoot(w http.ResponseWriter, r *http.Request) error {
	cert, err := LoadCertFromFile(s.config.CertFile)
	if err != nil {
		return err
	}
	curlPinnedPubKey := ""
	if bytes.Equal(cert.RawIssuer, cert.RawSubject) {
		if hash, err := CalculatePublicKeyHash(cert); err == nil {
			curlPinnedPubKey = EncodeCurlPinnedPublicKeyHash(hash)
		}
	}
	return webTemplate.Execute(w, &webTemplateConfig{
		Key:              s.config.Key,
		KeyDerivIter:     keyDerivIter,
		KeyLenBytes:      keyLenBytes,
		FileExpireAfter:  s.config.FileExpireAfter,
		CurlPinnedPubKey: curlPinnedPubKey,
	})
}

func (s *server) handleWebStatic(w http.ResponseWriter, r *http.Request) error {
	r.URL.Path = "/web" + r.URL.Path // This is a hack to get the embedded path
	http.FileServer(http.FS(webStaticFs)).ServeHTTP(w, r)
	return nil
}

func (s *server) handleClipboardGet(w http.ResponseWriter, r *http.Request) error {
	fields := r.Context().Value(routeCtx{}).([]string)
	file, err := s.getClipboardFile(fields[0])
	if err != nil {
		return errHTTPBadRequest
	}

	stat, err := os.Stat(file)
	if err != nil {
		return errHTTPNotFound
	}
	if stat.Mode()&os.ModeNamedPipe == 0 {
		w.Header().Set("Length", strconv.FormatInt(stat.Size(), 10))
	}
	f, err := os.Open(file)
	if err != nil {
		return errHTTPNotFound
	}
	defer f.Close()

	if _, err = io.Copy(w, f); err != nil {
		return err
	}
	if stat.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		os.Remove(file)
	}
	return nil
}

func (s *server) handleClipboardPut(w http.ResponseWriter, r *http.Request) error {
	fields := r.Context().Value(routeCtx{}).([]string)
	file, err := s.getClipboardFile(fields[0])
	if err != nil {
		return errHTTPBadRequest
	}

	// Check total file count limit (only if file didn't exist already)
	stat, _ := os.Stat(file)
	if stat == nil {
		if err := s.countLimiter.Add(1); err != nil {
			return errHTTPTooManyRequests
		}
	}

	// Make fifo device instead of file if type is set to "fifo"
	if r.Header.Get("Content-Type") == "inode/fifo" || r.URL.Query().Get("t") == "fifo" {
		os.Remove(file)
		if err := unix.Mkfifo(file, 0600); err != nil {
			s.countLimiter.Sub(1)
			return err
		}
	}

	// Create new file or truncate existing
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		s.countLimiter.Sub(1)
		return err
	}
	defer f.Close()
	defer s.updateStatsAndExpire()

	// Handle empty body
	if r.Body == nil {
		return errHTTPBadRequest
	}

	// Copy file contents (with file limit & total limit)
	fileSizeLimiter := newLimiter(s.config.FileSizeLimit)
	limitWriter := newLimitWriter(f, fileSizeLimiter, s.sizeLimiter)

	if _, err := io.Copy(limitWriter, r.Body); err != nil {
		os.Remove(file)
		if err == errLimitReached {
			return errHTTPBadRequest
		}
		return err
	}
	if err := r.Body.Close(); err != nil {
		os.Remove(file)
		return err
	}
	return nil
}

func (s *server) handleJoin(w http.ResponseWriter, r *http.Request) error {
	log.Printf("%s - %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	return joinTemplate.Execute(w, s.config)
}

func (s *server) getClipboardFile(file string) (string, error) {
	for _, reserved := range reservedFiles {
		if file == reserved {
			return "", errInvalidFileID
		}
	}
	return fmt.Sprintf("%s/%s", s.config.ClipboardDir, file), nil
}

func (s *server) auth(next handlerFnWithErr) handlerFnWithErr {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := s.authorize(r); err != nil {
			return err
		}
		return next(w, r)
	}
}

func (s *server) authorize(r *http.Request) error {
	if s.config.Key == nil {
		return nil
	}

	auth := r.Header.Get("Authorization")
	if encodedQueryAuth, ok := r.URL.Query()[authOverrideParam]; ok && len(encodedQueryAuth) > 0 {
		queryAuth, err := base64.StdEncoding.DecodeString(encodedQueryAuth[0])
		if err != nil {
			log.Printf("%s - %s %s - cannot decode query auth override", r.RemoteAddr, r.Method, r.RequestURI)
			return errHTTPUnauthorized
		}
		auth = string(queryAuth)
	}

	if m := authHmacRegex.FindStringSubmatch(auth); m != nil {
		return s.authorizeHmac(r, m)
	} else if m := authBasicRegex.FindStringSubmatch(auth); m != nil {
		return s.authorizeBasic(r, m)
	} else {
		log.Printf("%s - %s %s - auth header missing", r.RemoteAddr, r.Method, r.RequestURI)
		return errHTTPUnauthorized
	}
}

func (s *server) authorizeHmac(r *http.Request, matches []string) error {
	timestamp, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Printf("%s - %s %s - hmac timestamp conversion: %s", r.RemoteAddr, r.Method, r.RequestURI, err.Error())
		return errHTTPUnauthorized
	}

	ttlSecs, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Printf("%s - %s %s - hmac ttl conversion: %s", r.RemoteAddr, r.Method, r.RequestURI, err.Error())
		return errHTTPUnauthorized
	}

	hash, err := base64.StdEncoding.DecodeString(matches[3])
	if err != nil {
		log.Printf("%s - %s %s - hmac base64 conversion: %s", r.RemoteAddr, r.Method, r.RequestURI, err.Error())
		return errHTTPUnauthorized
	}

	// Recalculate HMAC
	data := []byte(fmt.Sprintf("%d:%d:%s:%s", timestamp, ttlSecs, r.Method, r.URL.Path))
	hm := hmac.New(sha256.New, s.config.Key.Bytes)
	if _, err := hm.Write(data); err != nil {
		log.Printf("%s - %s %s - hmac calculation: %s", r.RemoteAddr, r.Method, r.RequestURI, err.Error())
		return errHTTPUnauthorized
	}
	rehash := hm.Sum(nil)

	// Compare HMAC in constant time (to prevent timing attacks)
	if subtle.ConstantTimeCompare(hash, rehash) != 1 {
		log.Printf("%s - %s %s - hmac invalid", r.RemoteAddr, r.Method, r.RequestURI)
		return errHTTPUnauthorized
	}

	// Compare timestamp (to prevent replay attacks)
	maxAge := defaultMaxAuthAge
	if ttlSecs > 0 {
		maxAge = time.Second * time.Duration(ttlSecs)
	}
	if maxAge > 0 {
		age := time.Since(time.Unix(int64(timestamp), 0))
		if age > maxAge {
			log.Printf("%s - %s %s - hmac request age mismatch", r.RemoteAddr, r.Method, r.RequestURI)
			return errHTTPUnauthorized
		}
	}

	return nil
}

func (s *server) authorizeBasic(r *http.Request, matches []string) error {
	userPassBytes, err := base64.StdEncoding.DecodeString(matches[1])
	if err != nil {
		log.Printf("%s - %s %s - basic base64 conversion: %s", r.RemoteAddr, r.Method, r.RequestURI, err.Error())
		return errHTTPUnauthorized
	}

	userPassParts := strings.Split(string(userPassBytes), ":")
	if len(userPassParts) != 2 {
		log.Printf("%s - %s %s - basic invalid user/pass format", r.RemoteAddr, r.Method, r.RequestURI)
		return errHTTPUnauthorized
	}
	passwordBytes := []byte(userPassParts[1])

	// Compare HMAC in constant time (to prevent timing attacks)
	key := DeriveKey(passwordBytes, s.config.Key.Salt)
	if subtle.ConstantTimeCompare(key.Bytes, s.config.Key.Bytes) != 1 {
		log.Printf("%s - %s %s - basic invalid", r.RemoteAddr, r.Method, r.RequestURI)
		return errHTTPUnauthorized
	}

	return nil
}

func (s *server) serverManager() {
	ticker := time.NewTicker(managerTickerInterval)
	for {
		s.updateStatsAndExpire()
		<-ticker.C
	}
}

func (s *server) updateStatsAndExpire() {
	s.Lock()
	defer s.Unlock()

	// Expire visitors from rate limiter map
	for ip, v := range s.rateLimiter {
		if time.Since(v.lastSeen) > rateLimitExpungeAfter {
			delete(s.rateLimiter, ip)
		}
	}

	// Walk clipboard to update size/count limiters, and expire/delete files
	files, err := ioutil.ReadDir(s.config.ClipboardDir)
	if err != nil {
		log.Printf("error reading clipboard: %s", err.Error())
		return
	}
	numFiles := int64(0)
	totalSize := int64(0)
	for _, f := range files {
		if !s.maybeExpire(f) {
			numFiles++
			totalSize += f.Size()
		}
	}
	s.countLimiter.Set(numFiles)
	s.sizeLimiter.Set(totalSize)
	s.printStats()
}

func (s *server) printStats() {
	var countLimit, sizeLimit string
	if s.countLimiter.Limit() == 0 {
		countLimit = "no limit"
	} else {
		countLimit = fmt.Sprintf("max %d", s.countLimiter.Limit())
	}
	if s.sizeLimiter.Limit() == 0 {
		sizeLimit = "no limit"
	} else {
		sizeLimit = fmt.Sprintf("max %s", BytesToHuman(s.sizeLimiter.Limit()))
	}
	log.Printf("files: %d (%s), size: %s (%s), visitors: %d (last 3 minutes)",
		s.countLimiter.Value(), countLimit, BytesToHuman(s.sizeLimiter.Value()), sizeLimit, len(s.rateLimiter))
}

// maybeExpire deletes a file if it has expired and returns true if it did
func (s *server) maybeExpire(file os.FileInfo) bool {
	if s.config.FileExpireAfter == 0 || time.Since(file.ModTime()) <= s.config.FileExpireAfter {
		return false
	}
	if err := os.Remove(filepath.Join(s.config.ClipboardDir, file.Name())); err != nil {
		log.Printf("failed to remove clipboard entry after expiry: %s", err.Error())
	}
	log.Printf("removed expired entry: %s (%s)", file.Name(), BytesToHuman(file.Size()))
	return true
}

func (s *server) onlyIfWebUI(next handlerFnWithErr) handlerFnWithErr {
	return func(w http.ResponseWriter, r *http.Request) error {
		if !s.config.WebUI {
			return errHTTPBadRequest
		}

		return next(w, r)
	}
}

// limit wraps all HTTP endpoints and limits API use to a certain number of requests per second.
// This function was taken from https://www.alexedwards.net/blog/how-to-rate-limit-http-requests (MIT).
func (s *server) limit(next handlerFnWithErr) handlerFnWithErr {
	return func(w http.ResponseWriter, r *http.Request) error {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr // This should not happen in real life; only in tests.
		}

		limiter := s.getVisitorLimiter(ip)
		if !limiter.Allow() {
			return errHTTPTooManyRequests
		}

		return next(w, r)
	}
}

// getVisitorLimiter creates or retrieves a rate.Limiter for the given visitor.
// This function was taken from https://www.alexedwards.net/blog/how-to-rate-limit-http-requests (MIT).
func (s *server) getVisitorLimiter(ip string) *rate.Limiter {
	s.Lock()
	defer s.Unlock()

	v, exists := s.rateLimiter[ip]
	if !exists {
		limiter := rate.NewLimiter(rateLimitRequestsPerSecond, rateLimitBurstPerSecond)
		s.rateLimiter[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

func (s *server) fail(w http.ResponseWriter, r *http.Request, code int, err error) {
	log.Printf("%s - %s %s - %s", r.RemoteAddr, r.Method, r.RequestURI, err.Error())
	w.WriteHeader(code)
	w.Write([]byte(http.StatusText(code)))
}

type errHTTPNotOK struct {
	code   int
	status string
}

func (e errHTTPNotOK) Error() string {
	return fmt.Sprintf("http: %s", e.status)
}

var errListenAddrMissing = errors.New("listen address missing, add 'ListenAddr' to config or pass -listen")
var errKeyFileMissing = errors.New("private key file missing, add 'KeyFile' to config or pass -keyfile")
var errCertFileMissing = errors.New("certificate file missing, add 'CertFile' to config or pass -certfile")
var errClipboardDirNotWritable = errors.New("clipboard dir not writable by user")
var errInvalidFileID = errors.New("invalid file id")
var errNoMatchingRoute = errors.New("no matching route")
var errHTTPBadRequest = &errHTTPNotOK{http.StatusBadRequest, http.StatusText(http.StatusBadRequest)}
var errHTTPNotFound = &errHTTPNotOK{http.StatusNotFound, http.StatusText(http.StatusNotFound)}
var errHTTPTooManyRequests = &errHTTPNotOK{http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests)}
var errHTTPUnauthorized = &errHTTPNotOK{http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)}
