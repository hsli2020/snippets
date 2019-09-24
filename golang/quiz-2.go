package main

import "fmt"

func main() {
	quiz1()
	quiz2()
}

func quiz1() {
	i := 65
	fmt.Println(string(i))  // what is the output?
}

func quiz2() {
	var s1 []int
	var s2 = []int{}

	if s1 == nil {  // is this OK?
		fmt.Println("yes nil")
	} else {
		fmt.Println("nor nil")
	}

	if s2 == nil {  // is this OK?
		fmt.Println("yes nil")
	} else {
		fmt.Println("nor nil")
	}
}
