package main

import "fmt"

func fibonacci(c, quit chan int) {
	fmt.Println("fibonacci started")
	x, y := 0, 1
	for {
		fmt.Printf("x = %d\n", x)
		select {
		case c <- x:			
			x, y = y, x+y
		case quit_value := <-quit:
			fmt.Printf("quit_value = %d\n", quit_value)
			return
		}
	}
}

func main() {
	c := make(chan int)
	q := make(chan int)
	go func() {
		fmt.Println("Goroutine started")
		for i := 0; i < 5; i++ {
			value := <-c
			fmt.Printf("received %d\n", value)
		}
		q <- 999
	}()
	fmt.Println("calling fibonacci()")
	fibonacci(c, q)
}
