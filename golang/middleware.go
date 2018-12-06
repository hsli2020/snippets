package main

import (
    "io"
    "net/http"
)

const (
    MyAPIKey = "MY_EXAMPLE_KEY"
)

func main() {

    // Create an example endpoint/route
    http.Handle("/", Middleware(
        http.HandlerFunc(ExampleHandler),
        AuthMiddleware,
    ))

    // Run...
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

// Middleware (this function) makes adding more than one layer of middleware easy
// by specifying them as a list. It will run the last specified handler first.
func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
    for _, mw := range middleware {
        h = mw(h)
    }
    return h
}

// AuthMiddleware is an example of a middleware layer. It handles the request authorization
// by checking for a key in the url.
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        requestKey := r.URL.Query().Get("key")
        if len(requestKey) == 0 || requestKey != MyAPIKey {
            // Report Unauthorized
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusUnauthorized)
            io.WriteString(w, `{"error":"invalid_key"}`)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status":"ok"}`)
}
