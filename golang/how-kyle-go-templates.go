// https://gitlab.com/snippets/1662623

// Note that the templateFuncs refer to my apps helper package, left here in this example to help.
// Heavily inspired by the book https://www.sitepoint.com/premium/books/level-up-your-web-apps-with-go

package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Layout string

const (
	trimPrefix        = "templates/"
	Home       Layout = "home"
	Ajax       Layout = "ajax"
	User       Layout = "user"
)

var (
	layouts     *template.Template
	layoutFuncs = template.FuncMap{
		"yield": func() (string, error) {
			return "", fmt.Errorf("yield called unexpectedly.")
		},
		"appcss": func() string {
			stat, err := os.Stat("./static/css/app.css")
			if err != nil {
				return ""
			}
			return fmt.Sprintf("/static/css/app.css?%d", stat.ModTime().Unix())
		},
		"appjs": func() string {
			stat, err := os.Stat("./static/js/app.js")
			if err != nil {
				return ""
			}
			return fmt.Sprintf("/static/js/app.js?%d", stat.ModTime().Unix())
		},
	}

	templates     *template.Template
	templateFuncs = template.FuncMap{
		"lcfirst": helpers.LcFirst,
		"ucfirst": helpers.UcFirst,
	}

	errorTemplate = `
<html>
	<body>
		<h1>Error rendering template %s</h1>
		<p>%s</p>
	</body>
</html>
`
)

// We manually search and parse our templates over template.ParseGlob() so that we
// can guarantee unique template names using the file path, rather than the base name of a file.
func init() {
	// First parse the outer layouts.
	layouts = template.New("layouts").Funcs(layoutFuncs)
	lms, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	for _, lm := range lms {
		b, err := ioutil.ReadFile(lm)
		if err != nil {
			log.Fatal(err)
		}
		name := strings.TrimPrefix(lm, trimPrefix)
		template.Must(layouts.New(name).Parse(string(b)))
	}

	// Now find all templates that are children to layouts
	tp1, err := filepath.Glob("templates/*/*.html")
	if err != nil {
		log.Fatal(err)
	}

	tp2, err := filepath.Glob("templates/*/*/*.html")
	if err != nil {
		log.Fatal(err)
	}

	tms := []string{}
	tms = append(tms, tp1...)
	tms = append(tms, tp2...)

	// Now parse the templates.
	templates = template.New("templates").Funcs(templateFuncs)
	for _, tm := range tms {
		b, err := ioutil.ReadFile(tm)
		if err != nil {
			log.Fatal(err)
		}
		name := strings.TrimPrefix(tm, trimPrefix)
		template.Must(templates.New(name).Parse(string(b)))
	}
}

// GetLayout will return an appropriate layout to render a template.
func GetLayout(l Layout) *template.Template {
	return layouts.Lookup(fmt.Sprintf("%s%s.html", "layout_", string(l)))
}

// ByName will attempt to return a template instance by it's filename.
func ByName(name string) *template.Template {
	return templates.Lookup(name)
}

// Render will return the results of an html template.
func Render(l Layout, w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	data["CurrentUser"] = myuser.UserFromRequest(r)
	data["Flash"] = session.GetFlash(r, w)
	data["token"] = csrfbanana.Token(w, r, session.FromRequest(r))

	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()), err
		},
	}

	layoutClone, _ := GetLayout(l).Clone()
	layoutClone.Funcs(funcs)
	err := layoutClone.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errorTemplate, name, err.Error())
	}
}
