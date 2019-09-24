//== simple web server
package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, Gophers!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":3000", mux)
}

//== server struct

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

//== using server struct

package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gophers!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	httpServer := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	httpServer.ListenAndServe()
}

//== multiplexer

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

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about route"))
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home route"))
}
func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login route"))
}
func logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout route"))
}

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

//== handler

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

//== function as handler

// source : https://golang.org/src/net/http/server.go?s=61509:61556#L1993
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

//== http request

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

//==

package main

import (
	"net/http"
	"strconv"
)

func requestInspection(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("Method: "           + r.Method      + "\n")))
	w.Write([]byte(string("Protocol Version: " + r.Proto       + "\n")))
	w.Write([]byte(string("Host: "             + r.Host        + "\n")))
	w.Write([]byte(string("Referer: "          + r.Referer()   + "\n")))
	w.Write([]byte(string("User Agent: "       + r.UserAgent() + "\n")))
	w.Write([]byte(string("Remote Addr: "      + r.RemoteAddr  + "\n")))
	w.Write([]byte(string("Requested URL: "    + r.RequestURI  + "\n")))
	w.Write([]byte(string("Content Length: "   + strconv.FormatInt(r.ContentLength, 10) + "\n")))
	for key, value := range r.URL.Query() {
		w.Write([]byte(string("Query string: key=" + key + " value=" + value[0] + "\n")))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", requestInspection)

	http.ListenAndServe(":3000", mux)
}

//== http response

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

//==

package main

import (
	"net/http"
)

func unauthorized(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("you do not have permission to access this resource.\n"))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", unauthorized)

  http.ListenAndServe(":3000", mux)
}

//== http headers

type Header map[string][]string

//==

package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	hed := r.Header
	fmt.Fprintln(w, hed)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", headers)
	http.ListenAndServe(":3000", mux)
}

//== getting headers

	accept := r.Header.Get("Accept")

//== setting headers

package main

import "net/http"

func setHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ALLOWED", "GET, POST")
	w.Write([]byte("set allowed headers\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", setHeader)
	http.ListenAndServe(":3000", mux)
}

//== use of WriteHeader()

// WriteHeader is used to send http status code other than 200. Write calls 
// WriteHeader(http.StatusOK) before writing the data. It is important to call 
// WriteHeader before Write if status code we want to send is other than 200.

package main

import "net/http"

func setHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad request!\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", setHeader)
	http.ListenAndServe(":3000", mux)
}

//== query strings

package main

import "net/http"

func showQuery(w http.ResponseWriter, r *http.Request) {
	querystring := r.URL.Query()
	w.Write([]byte("query strings are\n"))
	w.Write([]byte("Name:" + querystring.Get("name") + "\n"))
	w.Write([]byte("Email:" + querystring.Get("email") + "\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", showQuery)
	http.ListenAndServe(":3000", mux)
}

//== working with cookies

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

//== setting cookie

//source : https://golang.org/src/net/http/cookie.go?s=3954:4002#L150
// SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers.
// The provided cookie must have a valid Name. Invalid cookies may be
// silently dropped.
func SetCookie(w ResponseWriter, cookie *Cookie) {
	if v := cookie.String(); v != "" {
		w.Header().Add("Set-Cookie", v)
	}
}

//==

package main

import (
	"net/http"
)

func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "cookie-1",
		Value: "hello world",
	}
	http.SetCookie(w, &cookie)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", setCookies)
	http.ListenAndServe(":3000", mux)
}

//== getting cookie

// source : https://github.com/golang/go/blob/go1.13.7/src/net/http/request.go#L408
// Cookies parses and returns the HTTP cookies sent with the request.
func (r *Request) Cookies() []*Cookie {
	return readCookies(r.Header, "")
}

In addition we also have Cookie method to get a single cookie by name.

//source : https://github.com/golang/go/blob/go1.13.7/src/net/http/request.go#L419
// Cookie returns the named cookie provided in the request or
// ErrNoCookie if not found.
// If multiple cookies match the given name, only one cookie will
// be returned.
func (r *Request) Cookie(name string) (*Cookie, error) {
	for _, c := range readCookies(r.Header, name) {
		return c, nil
	}
	return nil, ErrNoCookie
}

//==

package main

import (
	"fmt"
	"net/http"
)

func getCookies(w http.ResponseWriter, r *http.Request) {
	// get all cookies
	cookies := r.Cookies()
	for _, cookie := range cookies {
		fmt.Fprintln(w, cookie)
	}
	// get named cookie
	cookie, err := r.Cookie("cookie-1")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	fmt.Fprintln(w, cookie)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getCookies)
	http.ListenAndServe(":3000", mux)
}

//== working with sessions

//== get and set session data

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// for prod use secure key, not hard-coded
var store = sessions.NewCookieStore([]byte("1234"))

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "custom-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	val := session.Values["hello"]
	if val != nil {
		fmt.Fprintln(w, "retrieving existing session: ")
		fmt.Fprintln(w, val)
		return
	}
	session.Values["hello"] = "world"
	session.Save(r, w)
	fmt.Fprintln(w, "no existing session found, set value for session")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sessionHandler)
	http.ListenAndServe(":3000", mux)
}

//== flash messages

// Flash messages are convenience session data that can only be read once.
// Use case for flash messages could be one time confirmation message.

// gorilla toolkit has session.AddFlash() method that can be used to work with flash messages.

// Getting flash messages

	session, err := store.Get(r, "session-name")
	flashes := session.Flashes()

// Setting flash message

	session.AddFlash(message)
  
//==
