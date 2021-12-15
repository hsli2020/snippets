// https://github.com/lu4p/go-template-turbo-sample
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type templates struct {
	*template.Template
}

func (t templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.ExecuteTemplate(w, name, data)
}

func initTemplates() templates {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"toString": toString,
		"link":     link,
		"email": func() string {
			return "example@example.com"
		},
	})

	template.ParseGlob("templates/**.html") // ??

	t, err := t.ParseGlob("templates/**.html")
	if err != nil {
		log.Fatal(err)
	}

	return templates{t}
}

func main() {
	e := echo.New()

	e.Debug = true // TODO: this line should be removed in production
	e.Renderer = initTemplates()
	e.Use(middleware.Gzip(), middleware.Secure())
	// you should also add middleware.CSRF(), once you have forms

	e.GET("/", root)
	e.GET("/foo", foo)
	e.Static("/dist", "./dist")
	e.Start(":3000")
}

func root(c echo.Context) error {
	return c.Render(200, "index.html", map[string]interface{}{
		"title": "Root",
		"test":  "Hello, world!",
		"slice": []int{1, 2, 3},
	})
}

func foo(c echo.Context) error {
	return c.Render(200, "foo.html", map[string]interface{}{
		"title": "Foo",
	})
}

// toString converts any value to string
// functions that return a string are automatically escaped by html/template
func toString(v interface{}) string {
	return fmt.Sprint(v)
}

// link returns a styled "a" tag
// functions that return a template.HTML are not escaped,
//so all parameters need to be escaped to avoid xss
func link(location, name string) template.HTML {
	return escSprintf(`<a class="text-blue-600 no-underline hover:underline" href="%v">%v</a>`,
		location, name)
}

// escSprintf is like fmt.Sprintf but uses the escaped HTML equivalent of the args
func escSprintf(format string, args ...interface{}) template.HTML {
	for i, arg := range args {
		args[i] = template.HTMLEscapeString(fmt.Sprint(arg))
	}

	return template.HTML(fmt.Sprintf(format, args...))
}

/*
// ==================== templates/base.html => define head & foot

{{ define "head" }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.title}}</title>
    <link rel="stylesheet" href="/dist/main.css">
    <script src="/dist/main.js" defer></script>
</head>
<body>
    <div class="container mx-auto">
        <h1>{{.title}}</h1>
{{ end }}

{{ define "foot" }}
    </div>
</body>
</html>
{{ end }}

// ==================== templates/index.html

{{ template "head" . }}

{{ link "/foo" "Foo"}}
<p class="bg-gray-100">{{.test}}</p>
<p>{{toString .slice }}</p>

{{ template "foot" }}

// ==================== templates/foo.html

{{ template "head" . }}

{{ link "/" "Root Page"}}
<img class="h-32 mt-4" src="/dist/assets/gopher-go.png" alt="Image not loaded">

{{ template "foot" }}
*/
