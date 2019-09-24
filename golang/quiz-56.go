// Will the following program compile?
package main

import "fmt"

type worker interface {
	work() int
}

type person struct {
	worker
}

func main() {
	p := person{}
	fmt.Println(p)
}
