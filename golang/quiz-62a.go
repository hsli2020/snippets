package main

import (
	"fmt"
)

func main() {
	s := string([]rune{0xf8})
	fmt.Println(len(s))
}
