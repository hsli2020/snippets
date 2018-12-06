package main

import "fmt"

func main() {
	const N = 5

	for i := range [N]int{} { // this looks best, but memory
		fmt.Print(i)
	}
	fmt.Println("")
	for i := range [N]struct{}{} {
		fmt.Print(i)
	}
	fmt.Println("")
	for i := range [N][0]int{} {
		fmt.Print(i)
	}
	fmt.Println("")
	for i := range (*[N]int)(nil) {
		fmt.Print(i)
	}
	fmt.Println("")
}
