package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	//~ fmt.Println(iptr)
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("init", i)
	zeroval(i)
	fmt.Println("zeroval", i)
	
	zeroptr(&i)
	fmt.Println("zeroptr", i)
	
	fmt.Println("porter", &i)
}
