package main

import "fmt"

func main() {
	defer greet()()
	fmt.Println("vim-go")
}

func greet() func() {
	fmt.Println("Hello")
	return func() { fmt.Println("Bye") }
}
