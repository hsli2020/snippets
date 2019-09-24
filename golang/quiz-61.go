package main

import (
	"fmt"
)

func hello() int {
	j := 10
	defer func() {
		j = 15
	}()
	return j
}

func main() {
	fmt.Println(hello())
}
