package main

import (
    "fmt"
)

func hello() (j int) {
    j = 10
    defer func() {
      j = 15
    }()
    return
}

func main() {
    fmt.Println(hello())
}
