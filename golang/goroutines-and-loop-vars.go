package main

import (
	"fmt"
)

func main() {
	test1()
	fmt.Scanln()

	test2()
	fmt.Scanln()

	test3()
	fmt.Scanln()
}

// outputs: 3, 3, 3
func test1() {
	ints := []int{1, 2, 3}
	for _, i := range ints {
		go func() {
			fmt.Printf("Test1 %v\n", i)
		}()
	}
}

// solution 1: 1, 2, 3
func test2() {
	ints := []int{1, 2, 3}
	for _, i := range ints {
		go func(i int) {
			fmt.Printf("Test2 %v\n", i)
		}(i) // THIS IS KEY
	}
}

// solution 1: 1, 2, 3
func test3() {
	ints := []int{1, 2, 3}
	for _, i := range ints {
		i := i // THIS IS KEY
		go func() {
			fmt.Printf("Test3 %v\n", i)
		}()
	}
}
