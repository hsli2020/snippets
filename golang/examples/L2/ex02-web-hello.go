// Package hello is a simple App Engine application that replies to requests
// on /hello with a welcoming message.
package hello

import (
    "fmt"
    "net/http"
)

// init is run before the application starts serving.
func init() {
    // Handle all requests with path /hello with the helloHandler function.
    http.HandleFunc("/hello", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from the Go app")
}
