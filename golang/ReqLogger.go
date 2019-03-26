package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func SayGoodByeWorld(w http.ResponseWriter, r *http.Request) {
	html := "Good Bye World"
	w.Write([]byte(html))
}

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	html := "Hello World"
	w.Write([]byte(html))
}

func main() {
	fileName := "webrequests.log"

	// https://www.socketloop.com/tutorials/golang-how-to-save-log-messages-to-file
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	// direct all log messages to webrequests.log
	log.SetOutput(logFile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHelloWorld)
	mux.HandleFunc("/bye", SayGoodByeWorld)

	http.ListenAndServe(":8080", RequestLogger(mux))
}
