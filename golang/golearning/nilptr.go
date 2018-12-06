// Inspired by http://talks.golang.org/2014/go4java.slide#23

package main

import "fmt"
//import "github.com/davecgh/go-spew/spew"

type Person struct {
	name string
}

func (p *Person) Name() string {
	if p == nil {
		return "Anonymous"
	} else {
		return p.name
	}
}

func NewPerson(name string) *Person {
	var p = new(Person)
	p.name = name
	return p
}

func main() {
	var p1, p2 *Person

	p1 = NewPerson("Nemo")

	fmt.Println(p1.Name(), p2, p2.Name())
}
