// How to setup a nested HTML template in the Go Echo web framework
// https://blog.boatswain.io/post/setup-nested-html-template-in-go-echo-web-framework/
package main

import (
	"errors"
	"html/template"
	"io"
	"github.com/labstack/echo"
	"gitlab.com/ykyuen/golang-echo-template-example/handler"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	// Echo instance
	e := echo.New()

	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("view/home.html", "view/base.html"))
	templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))

	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Route => handler
	e.GET("/",      handler.HomeHandler)
	e.GET("/about", handler.AboutHandler)

	// Start the Echo server
	e.Logger.Fatal(e.Start(":1323"))
}

// ---------------------------------------------------------
// handler/about_handler.go

package handler

import (
	"net/http"
	"github.com/labstack/echo"
)

func AboutHandler(c echo.Context) error {
	// Please note the the second parameter "about.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"name": "About",
		"msg": "All about Boatswain!",
	})
}

// ---------------------------------------------------------
// view/base.html

{{define "base.html"}}
<!DOCTYPE html>
<html>
<head>
  <title>{{template "title" .}}</title>
</head>
<body>
  {{template "body" .}}
</body>
</html>
{{end}}

// ---------------------------------------------------------
// view/about.html

{{define "title"}}
    Boatswain Blog | {{index . "name"}}
{{end}}

{{define "body"}}
  <h1>{{index . "msg"}}</h1>
  <h2>This is the about page.</h2>
{{end}}
