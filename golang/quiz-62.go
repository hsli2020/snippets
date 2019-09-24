package main

import (
	"fmt"
)

func main() {
	i := [3]int{10, 10, 10}
	fmt.Println(i)
	j := []int(i[2:3])
	fmt.Println(j)
	i[2] = 900
	fmt.Println(j[0])
}
