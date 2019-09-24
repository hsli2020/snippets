package main

import (
  "fmt"
)

type Name1 struct {
  name string
}

type Name2 struct {
  name string
}

func main() {
  n1 := Name1{"Sam"}
  n2 := Name2(n1)
  fmt.Println(n2.name)
}
