package main

/*
 Channel and select {} illustrated by t9md
 ===========================================================
 sample program to understand goroutine communication via channel.

 based on go-tour's example
 https://tour.golang.org/concurrency/5
*/

import "fmt"

func fibonacci(out, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case out <- x: // if out-ch is available for write(means out-ch's buffer is not full), send next fib number x
			x, y = y, x+y
		case <-quit: // if any data come to quit-ch, then break inifinite loop by return.
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	/*
		main() start
		+-- main() --+
		|            |
		|            |
		|            |
		|            |
		+------------+
	*/
	out := make(chan int)
	quit := make(chan int)
	/*

		two channel (out, quit) created
		+-- main() --+
		|            |--------------
		|            |   out<int>
		|            |--------------
		|            |--------------
		|            |   quit<int>
		|            |--------------
		+------------+
	*/

	// Create anonymous func() and immediately execute as goroutine.
	// This function iterate 10 times of (read fib number from out-ch and Println).
	// After that, send `0` to quit-ch to notify other fibonacci() to break inifinit loop.
	// func() is closure referencing outer `out`, `quit` channel variable which originaly
	// created in enclosing main() function.
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-out) // pop data from out-ch 10 times
		}
		quit <- 0 // then send 0 to quit
	}()
	/*

		+--- main () -------------+              +---- go func() as goroutine ----------+
		|                         |              |                                      |
		|                         | ------------ |      for i := 0; i < 10; i++ {       | *1
		|                         |  out<int> ====>       fmt.Println(<-out) //(*1)     | Blocking here for now.
		|                         | ------------ |      }                               | since no data come from out-ch.
		|                         |              |                                      | First data will come after fibonacci() start
		|                         | ------------ |                                      |
		|                         |  quit<int> <== 0    quit <- 0                       | *2
		|                         | ------------ |                                      | Send 0 to quit-ch to break
		|                         |              |                                      | inifinite loop of fibonacci().
		+-------------------------+              +--------------------------------------+
	*/

	fibonacci(out, quit)
	/*
		Start fibonacci number generater fibonacci(), this fibonacci() function inifinitely generate fibonacci
		numbers, unless some data is sent to channel 'quit'
		Communication channel(out, quit) is passed by arguments.

		+-fibonacci(out, quit) in main()--+              +--- go func() as goroutine ------+
		|                                |              |                                 |
		|      x, y := 0, 1              |              |                                 |
		|      for {                     |              |                                 |
		|        select {                | ------------ |    for i := 0; i < 10; i++ {    | *1
		|        case out <- x:        x == out<int> ====>     fmt.Println(<-out) //(*1)  | groutine func() feel 10 times of data from out-ch
		|          x, y = x, x+y         | ------------ |    }                            | as sufficient.
		|                                |              |                                 | after for loop finish...
		|                                | ------------ |                                 |
		|        case <- quit:         <=== quit<int> <== 0  quit <- 0 (*2)               | *2
		|          fmt.Println("quit")   | ------------ |                                 | Send 0 to quit-ch for the purpose of stopping
		|          return                |              |                                 | inifinite loop of fibonacci().
		|        }                       |              |                                 |
		|      }                         |              |                                 |
		+--------------------------------+              +---------------------------------+
	*/

}