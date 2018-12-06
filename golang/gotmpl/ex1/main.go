package main

import (
	"net/http"
	"html/template"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Asit")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.ExecuteTemplate(w, "t2.html", "Golang")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/t1", handler1)
	http.HandleFunc("/t2", handler2)

	server.ListenAndServe()
}
