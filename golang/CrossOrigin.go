package main

import (
	"net/http"
	"strings"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func ReplyName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // usually, this line is enough, but you can add the following options.
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	URISegments := strings.Split(r.URL.Path, "/")

	if URISegments[2] != "" {
		w.Write([]byte(URISegments[2]))
	} else {
		w.Write([]byte("You gotta give a name for me to reply...."))
	}
}

func main() {
	// http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHelloWorld)
	mux.HandleFunc("/replyname/", ReplyName)

	http.ListenAndServe(":8080", mux)
}
