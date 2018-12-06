package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}
func main() {

	fmt.Println(person{"bob", 20})
	fmt.Println(&person{name:"Ann", age : 30})
	
	s := person{name: "sean", age: 50}
	fmt.Println(s.name)
	
	sp := &s
	fmt.Println(sp.age)
	
	sp.age = 51
	fmt.Println(sp.age)
}