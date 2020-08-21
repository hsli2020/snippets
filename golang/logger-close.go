package main

import (
    "io"
    "log"
    "os"
)

func SetupLogging() {
    logFile, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE, 0666)
    if err != nil {
        log.Panicln(err)
    }
    defer logFile.Close()	// NEVER work, File closed.

    log.SetOutput(io.MultiWriter(os.Stderr, logFile))
}

func main() {
    SetupLogging()
    log.Println("Test message")
}

/////////////////////////////////////////////////////////////

package main

import (
    "io"
    "log"
    "os"
)

func SetupLogging() *os.File {
    logFile, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE, 0666)
    if err != nil {
        log.Panicln(err)
    }

    log.SetOutput(io.MultiWriter(os.Stderr, logFile))
    return logFile	// return file handle
}

func main() {
    logf := SetupLogging()
    defer logf.Close()	// close file

    log.Println("Test message")
}

/////////////////////////////////////////////////////////////

package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func LogSetupAndDestruct() func() {
    logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    if err != nil {
        log.Panicln(err)
    }

    log.SetOutput(io.MultiWriter(os.Stderr, logFile))

    return func() {
        e := logFile.Close()
        if e != nil {
            fmt.Fprintf(os.Stderr, "Problem closing the log file: %s\n", e)
        }
    }
}

func main() {
    defer LogSetupAndDestruct()()

    log.Println("Test message")
}

/////////////////////////////////////////////////////////////
// Another option is to use runtime.SetFinalizer, but it's not always
// guaranteed to run before main exits.

func SetupLogging() {
    logFile, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE, 0666)
    if err != nil {
        log.Panicln(err)
    }
    runtime.SetFinalizer(logFile, func(h *os.File) {
        h.Close()
    })

    log.SetOutput(io.MultiWriter(os.Stderr, logFile))
}
