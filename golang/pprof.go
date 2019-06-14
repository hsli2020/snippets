package main

import (
    "fmt"
    "runtime"
    "sync"
    "github.com/davecheney/profile"
)

func main() {
    runtime.GOMAXPROCS(2)
    cfg := profile.Config{
        CPUProfile : true , 
        MemProfile:     true,
        ProfilePath:    ".",  // store profiles in current directory
        NoShutdownHook: true, // do not hook SIGINT
    }

    // p.Stop() must be called before the program exits to
    // ensure profiling information is written to disk.
    p := profile.Start(&cfg) 
    defer p.Stop()   

    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("Starting Go Routines")
    go func() {
        defer wg.Done()

        for char := 'a'; char < 'a'+26; char++ {
            fmt.Printf("%c ", char)
        }
    }()

    go func() {
        defer wg.Done()

        for number := 1; number < 27; number++ {
            fmt.Printf("%d ", number)
        }
    }()

    fmt.Println("Waiting To Finish")
    wg.Wait()

    fmt.Println("\nTerminating Program")
   
}
