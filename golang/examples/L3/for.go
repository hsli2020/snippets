package main

import "fmt"

func main() {
	
	i := 1
	for i <=3 {
		fmt.Println(i)
		i = i + 1
	}
	
	for j :=9; j <= 19; j++ {
		fmt.Println(j)
	}
	
	for {
		fmt.Println("loop")
		break
	}
}
