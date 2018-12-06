package main

import "fmt"
import "time"

func setTimeout(someFunc func(), milliseconds int) {

	timeout := time.Duration(milliseconds) * time.Millisecond

	// This spawns a goroutine and therefore does not block
	time.AfterFunc(timeout, someFunc)
}

func main() {

	printed := false

	print := func() {
		fmt.Println("This will print after x milliseconds")
		printed = true
	}

	// Make the timeout print after 5 seconds
	setTimeout(print, 5000)

	fmt.Println("This will print straight away")

	// Wait until it's printed our function string
	// before we close the program
	for {
		if printed {
			return
		}
	}
}
