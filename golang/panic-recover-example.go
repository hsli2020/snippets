package main

//import "fmt"

func main() {
	defer println("defer 1")

	level1()
}

func level1() {
	defer println("defer 3")
	defer func() {
		if err := recover(); err != nil {
			println("recovering in progress")
		}
	}()
	defer println("defer 2")

	level2()
}

func level2() {
	defer println("defer 4")
	panic("foo")
}

// defer 4
// defer 2
// recovering in progress
// defer 3
// defer 1
