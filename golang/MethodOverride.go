// https://www.alexedwards.net/blog/http-method-spoofing

package main

import (
    "net/http"
)

func MethodOverride(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Only act on POST requests.
        if r.Method == "POST" {

            // Look in the request body and headers for a spoofed method.
            // Prefer the value in the request body if they conflict.
            method := r.PostFormValue("_method")
            if method == "" {
                method = r.Header.Get("X-HTTP-Method-Override")
            }

            // Check that the spoofed method is a valid HTTP method and
            // update the request object accordingly.
            if method == "PUT" || method == "PATCH" || method == "DELETE" {
                r.Method = method
            }
        }

        // Call the next handler in the chain.
        next.ServeHTTP(w, r)
    })
}
