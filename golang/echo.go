Golang: Echo Protocol (RFC862)

Published: 2014-12-15

The Echo Protocol is one of the simplest possible internet protocols. The
server listens on port 7 and any data that is read is sent back to the client.
It can run either over TCP or UDP. In this article we'll see how to implement a
server and a client for the echo protocol in golang.

The Server (TCP)

First let's see the server over TCP. Create a file named echoservertcp.go, open
it in your favorite editor, and add the following lines:

package main

import (
       "log"
)

func main() {
     log.SetFlags(log.Lshortfile)  
}

We import the needed packages and then set the log flags to Lshortfile to
better spot errors when they happen. We then need to establish a TCP connection
on localhost port 7. Add net to the import list, and then add these lines to
main:

ln, err := net.Listen("tcp", ":7")
if err != nil {
    log.Fatal(err)
}

Now we need to listen for all the incoming connections and reply to each one.
We'll use a goroutine to process each connection so as to make our server
concurrent:

for {
    con, err := ln.Accept()
    if err != nil {
        log.Fatal(err)
    }
    go echo(con)
}

echo simply reads all the data from the client and sends it back to the client,
we do this with the handy io.Copy function (don't forget to import io):

func echo(con net.Conn) {
    _, err := io.Copy(con, con)
    if err != nil {
        log.Print(err)
    }
    err = con.Close()
    if err != nil {
        log.Print(err)
    }
}

We're done. Here is the final result (echoservertcp.go):

package main

import (
    "io"
    "log"
    "net"
)

func echo(con net.Conn) {
    _, err := io.Copy(con, con)
    if err != nil {
        log.Print(err)
    }
    err = con.Close()
    if err != nil {
        log.Print(err)
    }
}

func main() {
    log.SetFlags(log.Lshortfile)
    ln, err := net.Listen("tcp", ":7")
    if err != nil {
        log.Fatal(err)
    }
    for {
        con, err := ln.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go echo(con)
    }
    err = ln.Close()
    if err != nil {
        log.Fatal(err)
    }
}

The Client (TCP)

The client is simple as well. We:

    open a connection to localhost port 7
    read from stdin until EOF and send everything over the connection
    read the response and write it to stdout.

As simple as this is, there's one little gotcha. After writing the data, we
need to tell the server that we have nothing more to write, otherwise it will
block. So we check (trought type assertion) that our connection is effectivly a
TCP connection and if it is so we close the writing end. Here's the code
(echoclienttcp.go):

package main

import (
    "io"
    "log"
    "net"
    "os"
)

func main() {
    log.SetFlags(log.Lshortfile)
    con, err := net.Dial("tcp", ":7")
    if err != nil {
        log.Fatal(err)
    }
    _, err = io.Copy(con, os.Stdin)
    if err != nil {
        log.Fatal(err)
    }
    if tcpcon, ok := con.(*net.TCPConn); ok {
        tcpcon.CloseWrite()
    }
    _, err = io.Copy(os.Stdout, con)
    if err != nil {
        log.Fatal(err)
    }
    err = con.Close()
    if err != nil {
        log.Fatal(err)
    }
}

Let's try it!

Build the server and run it:

$ go build echoservertcp.go
$ ./echoservertcp
echoservertcp.go:21: listen tcp :7: bind: permission denied

We get an error because all ports below 1023 are reserved and can only be
accessed by root (see list of ports). So you'll need to login as root trought
su or sudo and rerun the command:

# ./echoservertcp

In another terminal build and run the client:

$ go build echoclienttcp.go
$ echo 'Hey, GNU rocks!' | ./echoclienttcp 
Hey, GNU rocks!

Hey, it works!
The Server (UDP)

In UDP we work with datagrams. We don't need to listen or accept any incoming
connection. We read a datagram and send back the response. In the TCP code we
block with net.Accept but here we'll use a semaphore to limit the number of
goroutines running. Here's the code (echoserverudp.go):

package main

import (
    "log"
    "net"
)

var sem = make(chan int, 100)

func echo(con net.PacketConn) {
    defer func() { <-sem }()
    buf := make([]byte, 4096)
    nr, addr, err := con.ReadFrom(buf)
    if err != nil {
        log.Print(err)
        return
    }
    nw, err := con.WriteTo(buf[:nr], addr)
    if err != nil {
        log.Print(err)
        return
    }
    if nw != nr {
        log.Printf("received %d bytes but sent %d\n", nr, nw)
    }
}

func main() {
    log.SetFlags(log.Lshortfile)
    con, err := net.ListenPacket("udp", ":7")
    if err != nil {
        log.Fatal(err)
    }
    for {
        sem <- 1
        go echo(con)
    }
    err = con.Close()
    if err != nil {
        log.Fatal(err)
    }
}

The Client (UDP)

In the client we:

    read data from stdin
    send it with a datagram
    read the response
    and write it to stdout

Code (echoclientudp.go):

package main

import (
    "io/ioutil"
    "log"
    "net"
    "os"
)

func main() {
    log.SetFlags(log.Lshortfile)
    con, err := net.Dial("udp", ":7")
    if err != nil {
        log.Fatal(err)
    }

    buf, err := ioutil.ReadAll(os.Stdin)

    nw, err := con.Write(buf)
    if err != nil {
        log.Fatal(err)
    }

    nr, err := con.Read(buf)
    if err != nil {
        log.Fatal(err)
    }

    if nr != nw {
        log.Fatalf("sent %d bytes but received %d\n", nw, nr)
    }

    _, err = os.Stdout.Write(buf[:nr])
    if err != nil {
        log.Fatal(err)
    }

    err = con.Close()
    if err != nil {
        log.Fatal(err)
    }
}

Please don't use the echo protocol. It's useless and what's more it's
dangerous: see CA-1996-01. It's appropriate only for learning purposes.
