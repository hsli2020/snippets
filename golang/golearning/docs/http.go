Package http
============
    func CanonicalHeaderKey(s string) string
    func DetectContentType(data []byte) string
    func Error(w ResponseWriter, error string, code int)
    func Handle(pattern string, handler Handler)
    func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    func ListenAndServe(addr string, handler Handler) error
    func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
    func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser
    func NotFound(w ResponseWriter, r *Request)
    func ParseHTTPVersion(vers string) (major, minor int, ok bool)
    func ParseTime(text string) (t time.Time, err error)
    func ProxyFromEnvironment(req *Request) (*url.URL, error)
    func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)
    func Redirect(w ResponseWriter, r *Request, urlStr string, code int)
    func Serve(l net.Listener, handler Handler) error
    func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
    func ServeFile(w ResponseWriter, r *Request, name string)
    func SetCookie(w ResponseWriter, cookie *Cookie)
    func StatusText(code int) string
    type Client
        func (c *Client) Do(req *Request) (*Response, error)
        func (c *Client) Get(url string) (resp *Response, err error)
        func (c *Client) Head(url string) (resp *Response, err error)
        func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
        func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
    type CloseNotifier
    type ConnState
        func (c ConnState) String() string
    type Cookie
        func (c *Cookie) String() string
    type CookieJar
    type Dir
        func (d Dir) Open(name string) (File, error)
    type File
    type FileSystem
    type Flusher
    type Handler
        func FileServer(root FileSystem) Handler
        func NotFoundHandler() Handler
        func RedirectHandler(url string, code int) Handler
        func StripPrefix(prefix string, h Handler) Handler
        func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
    type HandlerFunc
        func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
    type Header
        func (h Header) Add(key, value string)
        func (h Header) Del(key string)
        func (h Header) Get(key string) string
        func (h Header) Set(key, value string)
        func (h Header) Write(w io.Writer) error
        func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
    type Hijacker
    type ProtocolError
        func (err *ProtocolError) Error() string
    type Request
        func NewRequest(method, urlStr string, body io.Reader) (*Request, error)
        func ReadRequest(b *bufio.Reader) (*Request, error)
        func (r *Request) AddCookie(c *Cookie)
        func (r *Request) BasicAuth() (username, password string, ok bool)
        func (r *Request) Context() context.Context
        func (r *Request) Cookie(name string) (*Cookie, error)
        func (r *Request) Cookies() []*Cookie
        func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
        func (r *Request) FormValue(key string) string
        func (r *Request) MultipartReader() (*multipart.Reader, error)
        func (r *Request) ParseForm() error
        func (r *Request) ParseMultipartForm(maxMemory int64) error
        func (r *Request) PostFormValue(key string) string
        func (r *Request) ProtoAtLeast(major, minor int) bool
        func (r *Request) Referer() string
        func (r *Request) SetBasicAuth(username, password string)
        func (r *Request) UserAgent() string
        func (r *Request) WithContext(ctx context.Context) *Request
        func (r *Request) Write(w io.Writer) error
        func (r *Request) WriteProxy(w io.Writer) error
    type Response
        func Get(url string) (resp *Response, err error)
        func Head(url string) (resp *Response, err error)
        func Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
        func PostForm(url string, data url.Values) (resp *Response, err error)
        func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
        func (r *Response) Cookies() []*Cookie
        func (r *Response) Location() (*url.URL, error)
        func (r *Response) ProtoAtLeast(major, minor int) bool
        func (r *Response) Write(w io.Writer) error
    type ResponseWriter
    type RoundTripper
        func NewFileTransport(fs FileSystem) RoundTripper
    type ServeMux
        func NewServeMux() *ServeMux
        func (mux *ServeMux) Handle(pattern string, handler Handler)
        func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
        func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
        func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
    type Server
        func (srv *Server) ListenAndServe() error
        func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
        func (srv *Server) Serve(l net.Listener) error
        func (srv *Server) SetKeepAlivesEnabled(v bool)
    type Transport
        func (t *Transport) CancelRequest(req *Request)
        func (t *Transport) CloseIdleConnections()
        func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)
        func (t *Transport) RoundTrip(req *Request) (*Response, error)

Constants
=========

const (
        MethodGet     = "GET"
        MethodHead    = "HEAD"
        MethodPost    = "POST"
        MethodPut     = "PUT"
        MethodPatch   = "PATCH" // RFC 5789
        MethodDelete  = "DELETE"
        MethodConnect = "CONNECT"
        MethodOptions = "OPTIONS"
        MethodTrace   = "TRACE"
)

const (
        StatusContinue           = 100 // RFC 7231, 6.2.1
        StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
        StatusProcessing         = 102 // RFC 2518, 10.1

        StatusOK                   = 200 // RFC 7231, 6.3.1
        StatusCreated              = 201 // RFC 7231, 6.3.2
        StatusAccepted             = 202 // RFC 7231, 6.3.3
        StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
        StatusNoContent            = 204 // RFC 7231, 6.3.5
        StatusResetContent         = 205 // RFC 7231, 6.3.6
        StatusPartialContent       = 206 // RFC 7233, 4.1
        StatusMultiStatus          = 207 // RFC 4918, 11.1
        StatusAlreadyReported      = 208 // RFC 5842, 7.1
        StatusIMUsed               = 226 // RFC 3229, 10.4.1

        StatusMultipleChoices  = 300 // RFC 7231, 6.4.1
        StatusMovedPermanently = 301 // RFC 7231, 6.4.2
        StatusFound            = 302 // RFC 7231, 6.4.3
        StatusSeeOther         = 303 // RFC 7231, 6.4.4
        StatusNotModified      = 304 // RFC 7232, 4.1
        StatusUseProxy         = 305 // RFC 7231, 6.4.5

        StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
        StatusPermanentRedirect = 308 // RFC 7538, 3

        StatusBadRequest                   = 400 // RFC 7231, 6.5.1
        StatusUnauthorized                 = 401 // RFC 7235, 3.1
        StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
        StatusForbidden                    = 403 // RFC 7231, 6.5.3
        StatusNotFound                     = 404 // RFC 7231, 6.5.4
        StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
        StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
        StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
        StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
        StatusConflict                     = 409 // RFC 7231, 6.5.8
        StatusGone                         = 410 // RFC 7231, 6.5.9
        StatusLengthRequired               = 411 // RFC 7231, 6.5.10
        StatusPreconditionFailed           = 412 // RFC 7232, 4.2
        StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
        StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
        StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
        StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
        StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
        StatusTeapot                       = 418 // RFC 7168, 2.3.3
        StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
        StatusLocked                       = 423 // RFC 4918, 11.3
        StatusFailedDependency             = 424 // RFC 4918, 11.4
        StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
        StatusPreconditionRequired         = 428 // RFC 6585, 3
        StatusTooManyRequests              = 429 // RFC 6585, 4
        StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
        StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

        StatusInternalServerError           = 500 // RFC 7231, 6.6.1
        StatusNotImplemented                = 501 // RFC 7231, 6.6.2
        StatusBadGateway                    = 502 // RFC 7231, 6.6.3
        StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
        StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
        StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
        StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
        StatusInsufficientStorage           = 507 // RFC 4918, 11.5
        StatusLoopDetected                  = 508 // RFC 5842, 7.2
        StatusNotExtended                   = 510 // RFC 2774, 7
        StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)
