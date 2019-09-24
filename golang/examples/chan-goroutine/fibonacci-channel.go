package main

import "fmt"

// recursive
func fibor(n int) int {
	if n <= 1 {
		return n
	}
	return fibor(n-1) + fibor(n-2)
}

// iterative using closure
func fiboi() func() int {
	x, y := 0, 1
	return func() int {
		r := x
		x, y = y, x+y
		return r
	}
}

// iterative using channel
func fiboc(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	n := 10
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibor(i))
	}
	fmt.Println()
	next_fibo := fiboi()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", next_fibo())
	}
	fmt.Println()
	c := make(chan int, n)
	go fiboc(cap(c), c)
	for i := range c {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
