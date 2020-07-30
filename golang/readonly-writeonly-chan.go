package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go read(ch)
	write(10, ch)
}

func read(ch <-chan int) {
	// You can only receive messages from ch, you cannot send to it.
	i := <-ch
	fmt.Println(i)

	// this will NOT work.
	// ch <- 10
}

func write(i int, ch chan<- int) {
	// You can only send messages to ch, you cannot receive from it.
	ch <- i

	// this will NOT work.
	// i = <-ch
}
