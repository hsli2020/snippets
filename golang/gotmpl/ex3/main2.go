package main

// http://www.alexedwards.net/blog/serving-static-sites-with-go

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
