// https://www.alexedwards.net/blog/serving-static-sites-with-go

// $ touch main.go
// $ mkdir -p static/stylesheets
// $ touch static/example.html static/stylesheets/main.css
// $ mkdir templates
// $ touch templates/layout.html templates/example.html

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Print(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

// static/stylesheets/main.css
body {color: #c0392b}

// templates/layout.html

{{define "layout"}}
<!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<title>{{template "title"}}</title>
	<link rel="stylesheet" href="/static/stylesheets/main.css">
</head>
<body>
	{{template "body"}}
	<footer>Made with Go</footer>
</body>
</html>
{{end}}

// templates/example.html

{{define "title"}}A templated page{{end}}

{{define "body"}}
<h1>Hello from a templated page</h1>
{{end}}
