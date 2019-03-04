package main

import (
    "io"
    "fmt"
    "net/http"
)

const form = `
    <html><body>
        <form action="#" method="post" name="bar">
            <input type="text" name="in" />
            <input type="submit" value="submit"/>
        </form>
    </body></html>
`

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, request *http.Request) {
    io.WriteString(w, "<h1>hello, world</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    switch request.Method {
    case "GET":
        /* display the form to the user */
        io.WriteString(w, form)
    case "POST":
        /* handle the form data, note that ParseForm must
           be called before we can extract form data */
        //request.ParseForm();
        //io.WriteString(w, request.Form["in"][0])
        io.WriteString(w, request.FormValue("in"))
    }
}

func main() {
    fmt.Println("localhost:8088")

    http.HandleFunc("/test1", SimpleServer)
    http.HandleFunc("/test2", FormServer)
    if err := http.ListenAndServe(":8088", nil); err != nil {
        panic(err)
    }
}
