package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64 
}

func main() {
	originalHandler := http.HandlerFunc(postHandler)
	chain := New(filterContentType, setServerTimeCookie).Then(originalHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8000", nil)
}

// main handler 
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
		return
	}

	var tmpCity city
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tmpCity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}
	defer r.Body.Close()

	fmt.Printf("Got %s city with area of %d sq miles!\n",
	tmpCity.Name, tmpCity.Area)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Created"))
}

// middleware to check the content type header 
func filterContentType (handler http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			fmt.Println("Currently in the filterContentType middleware")
			// fmt.Println(r.Header)
			if r.Header.Get("Content-Type") != "application/json" {
				w.WriteHeader(http.StatusUnsupportedMediaType)
				w.Write([]byte("415 - Header Content-type missing"))
				return 
			}
			handler.ServeHTTP(w, r)
		})
}

// middleware that attaches the current time to the cookie before sending the response to the client
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cookie is a struct that represents an HTTP cookie as sent in the Set-Cookie
		// header of an HTTP response
		cookie := http.Cookie{
			Name: "Server-Time (UTC)",
			Value: strconv.Itoa(int(time.Now().Unix())),
		}
		http.SetCookie(w, &cookie)
		fmt.Println("Currently in setServerTimeCookie middleware")
		handler.ServeHTTP(w, r)
	})
}

type Middleware func(http.Handler) http.Handler
type Chain []Middleware 

// returns a Slice of middlewares 
func New(middlewares ...Middleware) Chain {
	var slice Chain
	return append(slice, middlewares...)
}

func (c Chain) Then(originalHandler http.Handler) http.Handler {
	if originalHandler == nil {
		originalHandler = http.DefaultServeMux
	}

	for i := range c {
		// Equivalent to m1(m2(m3(originalHandler)))
		originalHandler = c[len(c) -1 -i](originalHandler)
	}
	return originalHandler
}
