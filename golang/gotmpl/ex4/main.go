package main

import (
	"net/http"

	"calhoun.io/views"
)

var index *views.View
var contact *views.View

func main() {
	index = views.NewView("bootstrap", "views/index.gohtml")
	contact = views.NewView("bootstrap", "views/contact.gohtml")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Render(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	contact.Render(w, nil)
}
