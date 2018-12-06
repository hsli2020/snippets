package main

// golang pop quiz: does this program terninate?

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Printf("%+v", v)
}
