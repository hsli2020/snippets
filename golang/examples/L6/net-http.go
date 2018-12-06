ListenAndServe starts an HTTP server with a given address and handler.
The handler is usually nil, which means to use DefaultServeMux.
Handle and HandleFunc add handlers to DefaultServeMux:

// func Handle(pattern string, handler Handler)
http.Handle("/foo", fooHandler)

// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

// func ListenAndServe(addr string, handler Handler) error
log.Fatal(http.ListenAndServe(":8080", nil))

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)
    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against
a list of registered patterns and calls the handler for the pattern that most closely matches
the URL.

Patterns name fixed, rooted paths, like "/favicon.ico", or rooted subtrees, like "/images/"
(note the trailing slash). Longer patterns take precedence over shorter ones, so that if there
are handlers registered for both "/images/" and "/images/thumbnails/", the latter handler will
be called for paths beginning "/images/thumbnails/" and the former will receive requests for
any other paths in the "/images/" subtree.

Note that since a pattern ending in a slash names a rooted subtree, the pattern "/" matches
all paths not matched by other registered patterns, not just the URL with Path == "/".

Patterns may optionally begin with a host name, restricting matches to URLs on that host only.
Host-specific patterns take precedence over general patterns, so that a handler might register
for the two patterns "/codesearch" and "codesearch.google.com/" without also taking over requests
for "http://www.google.com/".

ServeMux also takes care of sanitizing the URL request path, redirecting any request containing .
or .. elements to an equivalent .- and ..-free URL.

type ServeMux struct {
    // contains filtered or unexported fields
}
    func NewServeMux() *ServeMux
    func (mux *ServeMux) Handle(pattern string, handler Handler)
    func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
    func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)

type Request
    func NewRequest(method, urlStr string, body io.Reader) (*Request, error)
    func ReadRequest(b *bufio.Reader) (req *Request, err error)
    func (r *Request) AddCookie(c *Cookie)
    func (r *Request) BasicAuth() (username, password string, ok bool)
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
    func (r *Request) Write(w io.Writer) error
    func (r *Request) WriteProxy(w io.Writer) error

type Response
    func (r *Response) Cookies() []*Cookie
    func (r *Response) Location() (*url.URL, error)
    func (r *Response) ProtoAtLeast(major, minor int) bool
    func (r *Response) Write(w io.Writer) error

func Get(url string) (resp *Response, err error)
func Head(url string) (resp *Response, err error)
func Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
func PostForm(url string, data url.Values) (resp *Response, err error)
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)

type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(int)
}

func Error(w ResponseWriter, error string, code int)
func NotFound(w ResponseWriter, r *Request)
func Redirect(w ResponseWriter, r *Request, urlStr string, code int)
func ServeFile(w ResponseWriter, r *Request, name string)
func SetCookie(w ResponseWriter, cookie *Cookie)
func StatusText(code int) string

func FileServer(root FileSystem) Handler
func NotFoundHandler() Handler
func RedirectHandler(url string, code int) Handler
func StripPrefix(prefix string, h Handler) Handler
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler

