package main

import (
    "fmt"
)

func pr(v ...interface{}) {
    fmt.Printf("%v\n", v...)
    fmt.Printf("%+v\n", v...)
    fmt.Printf("%#v\n", v...)
}

func main() {
    a := []string{"1", "2", "3", "4 5"}
    pr(a)
}
