package main

import "fmt"
import "time"

func main() {
	fmt.Println(time.Now().Weekday())
	
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("It's before noon")
		default:
			fmt.Println("after")
	}
}