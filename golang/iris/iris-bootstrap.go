File: structuring/bootstrap/bootstrap/bootstrapper.go

package bootstrap

import (
	"time"

	"github.com/gorilla/securecookie"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time

	Sessions *sessions.Sessions
}

// New returns a new Bootstrapper.
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// SetupViews loads the templates.
func (b *Bootstrapper) SetupViews(viewsDir string) {
	b.RegisterView(iris.HTML(viewsDir, ".html").Layout("shared/layout.html"))
}

// SetupSessions initializes the sessions, optionally.
func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.AppName,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

// SetupWebsockets prepares the websocket server.
func (b *Bootstrapper) SetupWebsockets(endpoint string, onConnection websocket.ConnectionFunc) {
	ws := websocket.New(websocket.Config{})
	ws.OnConnection(onConnection)

	b.Get(endpoint, ws.Handler())
	b.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}

// SetupErrorHandlers prepares the http error handlers
// `(context.StatusCodeNotSuccessful`,  which defaults to < 200 || >= 400 but you can change it).
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./views")
	b.SetupSessions(24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}

File: structuring/bootstrap/main.go

package main

import (
	"github.com/kataras/iris/_examples/structuring/bootstrap/bootstrap"
	"github.com/kataras/iris/_examples/structuring/bootstrap/middleware/identity"
	"github.com/kataras/iris/_examples/structuring/bootstrap/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Awesome App", "kataras2006@hotmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}

File: structuring/bootstrap/main_test.go

package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

// go test -v
func TestApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app.Application)

	// test our routes
	e.GET("/").Expect().Status(httptest.StatusOK)
	e.GET("/follower/42").Expect().Status(httptest.StatusOK).
		Body().Equal("from /follower/{id:long} with ID: 42")
	e.GET("/following/52").Expect().Status(httptest.StatusOK).
		Body().Equal("from /following/{id:long} with ID: 52")
	e.GET("/like/64").Expect().Status(httptest.StatusOK).
		Body().Equal("from /like/{id:long} with ID: 64")

	// test not found
	e.GET("/notfound").Expect().Status(httptest.StatusNotFound)
	expectedErr := map[string]interface{}{
		"app":     app.AppName,
		"status":  httptest.StatusNotFound,
		"message": "",
	}
	e.GET("/anotfoundwithjson").WithQuery("json", nil).
		Expect().Status(httptest.StatusNotFound).JSON().Equal(expectedErr)
}

File: structuring/bootstrap/middleware/identity/identity.go

package identity

import (
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/_examples/structuring/bootstrap/bootstrap"
)

// New returns a new handler which adds some headers and view data
// describing the application, i.e the owner, the startup time.
func New(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		// response headers
		ctx.Header("App-Name", b.AppName)
		ctx.Header("App-Owner", b.AppOwner)
		ctx.Header("App-Since", time.Since(b.AppSpawnDate).String())

		ctx.Header("Server", "Iris: https://iris-go.com")

		// view data if ctx.View or c.Tmpl = "$page.html" will be called next.
		ctx.ViewData("AppName", b.AppName)
		ctx.ViewData("AppOwner", b.AppOwner)
		ctx.Next()
	}
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	h := New(b)
	b.UseGlobal(h)
}

File: structuring/bootstrap/routes/follower.go

package routes

import (
	"github.com/kataras/iris"
)

// GetFollowerHandler handles the GET: /follower/{id}
func GetFollowerHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}

File: structuring/bootstrap/routes/following.go

package routes

import (
	"github.com/kataras/iris"
)

// GetFollowingHandler handles the GET: /following/{id}
func GetFollowingHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}

File: structuring/bootstrap/routes/index.go

package routes

import (
	"github.com/kataras/iris"
)

// GetIndexHandler handles the GET: /
func GetIndexHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Index Page")
	ctx.View("index.html")
}

File: structuring/bootstrap/routes/like.go

package routes

import (
	"github.com/kataras/iris"
)

// GetLikeHandler handles the GET: /like/{id}
func GetLikeHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}

File: structuring/bootstrap/routes/routes.go

package routes

import (
	"github.com/kataras/iris/_examples/structuring/bootstrap/bootstrap"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	b.Get("/", GetIndexHandler)
	b.Get("/follower/{id:long}", GetFollowerHandler)
	b.Get("/following/{id:long}", GetFollowingHandler)
	b.Get("/like/{id:long}", GetLikeHandler)
}
