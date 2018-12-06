package main

import     "fmt"
//import   "math"
//import   "math/rand"
//import   "time"
//import   "strconv"

type Foo int64

func main() {
    fmt.Println("Let's Go!")

}

/*
////////////////////////////////////////////////////////////
package main

import (
    "encoding/gob"
    "fmt"
    "net"
)

func server() {
    // listen on a port
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        // accept a connection
        c, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        // handle the connection
        go handleServerConnection(c)
    }
}

func handleServerConnection(c net.Conn) {
    // receive the message
    var msg string
    err := gob.NewDecoder(c).Decode(&msg)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Received", msg)
    }
    c.Close()
}

func client() {
    // connect to the server
    c, err := net.Dial("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    // send the message
    msg := "Hello World"
    fmt.Println("Sending", msg)
    err = gob.NewEncoder(c).Encode(msg)
    if err != nil {
        fmt.Println(err)
    }
    c.Close()
}

func main() {
     go server()
     go client()

     var input string
     fmt.Scanln(&input)
}
////////////////////////////////////////////////////////////
package main

import ("net/http" ; "io")

func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
          "Content-Type",
          "text/html",
    )
    io.WriteString(res,
`<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`,
    )
}

func main() {
    // http.Handle("/assets/",
    //     http.StripPrefix(
    //         "/assets/",
    //         http.FileServer(http.Dir("assets")),
    //     ),
    // )
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":9000", nil)
}
////////////////////////////////////////////////////////////
package main

import (
    "fmt"
    "net"
    "net/rpc"
)

type Server struct {}

func (this *Server) Negate(i int64, reply *int64) error {
    *reply = -i
    return nil
}

func server() {
    rpc.Register(new(Server))
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        c, err := ln.Accept()
        if err != nil {
            continue
        }
        go rpc.ServeConn(c)
    }
}

func client() {
    c, err := rpc.Dial("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    var result int64
    err = c.Call("Server.Negate", int64(999), &result)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Server.Negate(999) =", result)
    }
}

func main() {
    go server()
    go client()
    var input string
    fmt.Scanln(&input)
}
////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////
*/
