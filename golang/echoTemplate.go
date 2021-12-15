package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/dustin/go-humanize"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("libra"))))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(Auth())
	e.HTTPErrorHandler = errorHandler
	e.HideBanner = true

	funcMap := template.FuncMap{
		"formatNum": func(num uint64) string {
			return humanize.Comma(int64(num))
		},
		"formatLibraNum": func(num uint64) string {
			return humanize.Comma(int64(num / 1e6))
		},
	}
	t := &Template{
		templates: template.Must(
            template.New("main").Funcs(funcMap).ParseGlob("template/*.html")),
	}
	e.Renderer = t

	e.Static("/", "assets")

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", "")
	})

	e.GET("/mint", func(c echo.Context) error {
		accAddr, _ := getAccount(c.Get("walletm").(string))
		return c.Render(http.StatusOK, "mint.html", accAddr)
	})

	e.Logger.Fatal(e.Start(*s))
}
