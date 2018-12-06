// server.go
package server

import "log"

// New creates a new server that accepts Work requests
// through the req channel.
func New() (req chan<- *Work) {
    wc := make(chan *Work)
    go serve(wc)
    return wc
}

type Work struct {
    Op    func(int, int) int
    A, B  int
    Reply chan int  // Server sends result on this channel.
}

func serve(wc <-chan *Work) {
    for w := range wc {
        go safelyDo(w)
    }
}

func safelyDo(w *Work) {
    // Regain control of panicking goroutine to avoid
    // killing the other executing goroutines.
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(w)
}

func do(w *Work) {
    w.Reply <- w.Op(w.A, w.B)
}

// example_test.go
package server_test

import (
    server "."
    "fmt"
    "time"
)

func main() {
    s := server.New()

    divideByZero := &server.Work{
        Op:    func(a, b int) int { return a / b },
        A:     100,
        B:     0,
        Reply: make(chan int),
    }
    s <- divideByZero

    select {
    case res := <-divideByZero.Reply:
        fmt.Println(res)
    case <-time.After(time.Second):
        fmt.Println("No result in one second.")
    }
    // Output: No result in one second.
}
