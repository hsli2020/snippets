package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "views/layouts"
var index *template.Template
var contact *template.Template

func main() {
	var err error
	files := append(layoutFiles(), "views/index.gohtml")
	index, err = template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	files = append(layoutFiles(), "views/contact.gohtml")
	contact, err = template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.ExecuteTemplate(w, "bootstrap", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	contact.ExecuteTemplate(w, "bootstrap", nil)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}
