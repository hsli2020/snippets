package main

// http://play.golang.org/p/ZeJN62RUWY

import (
	"fmt"
)

// define an animal interface
type Animal interface {
	Sound() // all animals makes sounds
	Move()  // and they are moving
}

// An abstract Dog satisfies the Animal interface
type AbstractDog struct{}

func (dog AbstractDog) Sound() {
	fmt.Println("Woof Woof!")
}
func (dog AbstractDog) Move() {
	fmt.Println("Walk walk")
}

// My little doggy
type Fido struct {
	AbstractDog
}
func (myDog Fido) Sound() {
	fmt.Println("yap yap yap!")
}

// Big Bad Guard Dog
type Zeus struct {
	AbstractDog
}

func (guard Zeus) Sound() {
	fmt.Println("GRRRRR GRRRRR bite bite GRRRR")
}

func NewDog() Animal {
	return Fido{}
}

func NewGuardDog() Animal {
	return Zeus{}
}

func main() {
	dog := NewDog()
	dog.Sound()
	dog.Move()
	fmt.Println("")
	dog = NewGuardDog()
	dog.Sound()
	dog.Move()
}
