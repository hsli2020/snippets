echo_server_tcp.go

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

echo_client_tcp.go

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

echo_server_udp.go

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

echo_client_udp.go

package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)
	con, err := net.Dial("udp", "xojoc.pw:7")
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
