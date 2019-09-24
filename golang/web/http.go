//## routes.go
package main

import "net/http"

var mux = http.NewServeMux()

func registerRoutes() {
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
}

//## handlers.go
package main

import "net/http"

func about(w http.ResponseWriter, r *http.Request)  { w.Write([]byte("about route"))  }
func home(w http.ResponseWriter, r *http.Request)   { w.Write([]byte("home route"))   }
func login(w http.ResponseWriter, r *http.Request)  { w.Write([]byte("login route"))  }
func logout(w http.ResponseWriter, r *http.Request) { w.Write([]byte("logout route")) }

//## main.go
package main

import "net/http"

func main() {
	registerRoutes()
	httpServer := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	httpServer.ListenAndServe()
}

// from https://golang.org/src/net/http/server.go?s=77156:81268#L2480
// A Server defines parameters for running an HTTP server.

// The zero value for Server is a valid configuration.
type Server struct {
	Addr              string  // TCP address to listen on, ":http" if empty
	Handler           Handler // handler to invoke, http.DefaultServeMux if nil
	TLSConfig         *tls.Config
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
	TLSNextProto      map[string]func(*Server, *tls.Conn, Handler)
	ConnState         func(net.Conn, ConnState)
	ErrorLog          *log.Logger
	BaseContext       func(net.Listener) context.Context
	ConnContext       func(ctx context.Context, c net.Conn) context.Context
	disableKeepAlives int32     // accessed atomically.
	inShutdown        int32     // accessed atomically (non-zero means we're in Shutdown)
	nextProtoOnce     sync.Once // guards setupHTTP2_* init
	nextProtoErr      error     // result of http2.ConfigureServer if used
	mu                sync.Mutex
	listeners         map[*net.Listener]struct{}
	activeConn        map[*conn]struct{}
	doneChan          chan struct{}
	onShutdown        []func()
}

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

// function as handler

// source : https://golang.org/src/net/http/server.go?s=61509:61556#L1993
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

// source : https://golang.org/src/net/http/request.go?s=3252:11812#L97
type Request struct {
	Method           string
	URL              *url.URL
	Proto            string // "HTTP/1.0"
	ProtoMajor       int    // 1
	ProtoMinor       int    // 0
	Header           Header
	Body             io.ReadCloser
	GetBody          func() (io.ReadCloser, error)
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Host             string
	Form             url.Values
	PostForm         url.Values
	MultipartForm    *multipart.Form
	Trailer          Header
	RemoteAddr       string
	RequestURI       string
	TLS              *tls.ConnectionState
	Cancel           <-chan struct{}
	Response         *Response
	ctx              context.Context
}

// source : https://golang.org/src/net/http/response.go?s=731:4298#L25
type Response struct {
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Proto            string // e.g. "HTTP/1.0"
	ProtoMajor       int    // e.g. 1
	ProtoMinor       int    // e.g. 0
	Header           Header
	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Uncompressed     bool
	Trailer          Header
	Request          *Request
	TLS              *tls.ConnectionState
}

// We do not directly work with Response struct, instead we can build the http response 
// using ResponseWriter interface. ResponseWriter interface is defined as

// source : https://golang.org/src/net/http/server.go?s=2985:5848#L84
type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

type Header map[string][]string

// accept := r.Header.Get("Accept")

// source : https://github.com/golang/go/blob/go1.13.7/src/net/http/cookie.go#L19
type Cookie struct {
	Name       string
	Value      string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge     int
	Secure     bool
	HttpOnly   bool
	SameSite   SameSite
	Raw        string
	Unparsed   []string // Raw text of unparsed attribute-value pairs
}

//source : https://golang.org/src/net/http/cookie.go?s=3954:4002#L150
// SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers.
// The provided cookie must have a valid Name. Invalid cookies may be silently dropped.
func SetCookie(w ResponseWriter, cookie *Cookie) {
	if v := cookie.String(); v != "" {
		w.Header().Add("Set-Cookie", v)
	}
}

func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "cookie-1",
		Value: "hello world",
	}
	http.SetCookie(w, &cookie)
}

// source : https://github.com/golang/go/blob/go1.13.7/src/net/http/request.go#L408
// Cookies parses and returns the HTTP cookies sent with the request.
func (r *Request) Cookies() []*Cookie {
	return readCookies(r.Header, "")
}
