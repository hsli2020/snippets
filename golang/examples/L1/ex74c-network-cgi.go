package main

import (
    "net/http/cgi"
    "net/http"
    "fmt"
)

func errorResponse(code int, msg string) {
    fmt.Printf("Status:%d %s\r\n", code, msg)
    fmt.Printf("Content-Type: text/plain\r\n")
    fmt.Printf("\r\n")
    fmt.Printf("%s\r\n", msg)
}

func main() {

    var req *http.Request
    var err error
    req, err = cgi.Request()
    if err != nil {
        errorResponse(500, "cannot get cgi request" + err.Error())
        return
    }

    // Use req to handle request

    fmt.Printf("Content-Type: text/plain\r\n")
    fmt.Printf("\r\n")
    fmt.Printf("req=%v\r\n", req)
}
