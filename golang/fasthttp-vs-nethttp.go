

net/http -> fasthttp conversion table:

    All the pseudocode below assumes w, r and ctx have these types:

  var (
  	w http.ResponseWriter
  	r *http.Request
  	ctx *fasthttp.RequestCtx
  )

    r.Body -> ctx.PostBody()
    r.URL.Path -> ctx.Path()
    r.URL -> ctx.URI()
    r.Method -> ctx.Method()
    r.Header -> ctx.Request.Header
    r.Header.Get() -> ctx.Request.Header.Peek()
    r.Host -> ctx.Host()
    r.Form -> ctx.QueryArgs() + ctx.PostArgs()
    r.PostForm -> ctx.PostArgs()
    r.FormValue() -> ctx.FormValue()
    r.FormFile() -> ctx.FormFile()
    r.MultipartForm -> ctx.MultipartForm()
    r.RemoteAddr -> ctx.RemoteAddr()
    r.RequestURI -> ctx.RequestURI()
    r.TLS -> ctx.IsTLS()
    r.Cookie() -> ctx.Request.Header.Cookie()
    r.Referer() -> ctx.Referer()
    r.UserAgent() -> ctx.UserAgent()
    w.Header() -> ctx.Response.Header
    w.Header().Set() -> ctx.Response.Header.Set()
    w.Header().Set("Content-Type") -> ctx.SetContentType()
    w.Header().Set("Set-Cookie") -> ctx.Response.Header.SetCookie()
    w.Write() -> ctx.Write(), ctx.SetBody(), ctx.SetBodyStream(), ctx.SetBodyStreamWriter()
    w.WriteHeader() -> ctx.SetStatusCode()
    w.(http.Hijacker).Hijack() -> ctx.Hijack()
    http.Error() -> ctx.Error()
    http.FileServer() -> fasthttp.FSHandler(), fasthttp.FS
    http.ServeFile() -> fasthttp.ServeFile()
    http.Redirect() -> ctx.Redirect()
    http.NotFound() -> ctx.NotFound()
    http.StripPrefix() -> fasthttp.PathRewriteFunc

