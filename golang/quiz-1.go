package main

import "fmt"

func quiz1() {
	var a, b int
	var c = &b
	switch *c {
	case a:
		fmt.Println("a")
	case b:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}

func quiz2() {
	var a, b *int
	var c = b
	switch c {
	case a:
		fmt.Println("a")
	case b:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}

func quiz3() {
	var a, b *int
	var c = b
	switch true {
	case nil == a:
		fmt.Println("a")
	case b == c:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}

func quiz4a() {
	const (
		a = 1
		b
		c
	)
	fmt.Println(b)
}

func quiz4() {
	const (
		a = 0
		b
	)
	fmt.Println(b)

	/*var c int
	switch c {
	case a:
		fmt.Println("a")
	case 1:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}*/
}

func main() {
	//quiz1()
	//quiz2()
	//quiz3()
	quiz4a()
	quiz4()
}
