package main

// https://twitter.com/val_deleplace/status/1476907441459798055

import "fmt"

// What is output?
func main() {
	var x = []string{"A", "B", "C"}

	for i, s := range x {
		fmt.Print(i, s, " ")
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}
