package main

// https://twitter.com/val_deleplace/status/1476239147601997828

import "fmt"

func main() {
	var x []string
	x = append(x, "a")
	x = append(x, "b")
	x = append(x, "c")

	for i, s := range x {
		fmt.Print(i, s, " ")
		x = append(x, "z")
		x[i+1] = "z"
	}
}
