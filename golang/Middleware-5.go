package main

import(
    "fmt"
    "time"
)

type MyContext struct {
    handlers[] func(c * MyContext)
    index int8
}

func(c * MyContext) Use(f func(c * MyContext)) {
    c.handlers = append(c.handlers, f)
}

func(c * MyContext) Next() {
    c.index++
    c.handlers[c.index](c)
}

func(c * MyContext) GET(path string, f func(c * MyContext)) {
    c.handlers = append(c.handlers, f)
}

func(c * MyContext) Run() {
    c.handlers[0](c)
}

// Authentication middleware
func AuthMiddleware(c * MyContext) func(c * MyContext) {
    return func(c * MyContext) {
        fmt.Println("[AuthMiddleware Start]")
        c.Next()
    }
}

// log middleware
func LogMiddleware(c * MyContext) func(c * MyContext) {
    return func(c * MyContext) {
        start: = time.Now()
        fmt.Printf("[LogHandle] start at: %d \n", start.Unix())
        c.Next()
        fmt.Printf("cost %f second \n", time.Since(start).Seconds())
    }
}

func main() {
    c: = & MyContext {}

    // Add middleware
    c.Use(AuthMiddleware(c))
    c.Use(LogMiddleware(c))

    // Simulate common routing
    c.GET("/", func(c * MyContext) {
        fmt.Println("hello go")
    })

    // run
    c.Run()
}