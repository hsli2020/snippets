package main

import "fmt"

type AnimalType int
type AnimalCreator func() Animal

const (
	DOG	AnimalType	= iota
	CAT
	LIZARD
	JAVA_PROGRAMMER
)

type AnimalFactory struct {
	m map[AnimalType]AnimalCreator
}

func NewAnimalFactory() *AnimalFactory {
	return &AnimalFactory{m: make(map[AnimalType]AnimalCreator, 10)}
}

func (a *AnimalFactory) MakeAnimal(t AnimalType) Animal {
	fn, ok := a.m[t]
	if !ok {
		panic("whatever")
	}
	return fn()
}

func (a *AnimalFactory) Register(t AnimalType, fn AnimalCreator) {
	a.m[t] = fn
}

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d *Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Lizard struct {
}

func (l *Lizard) Speak() string {
	return "????"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) Speak() string {
	return "something about design patterns"
}

func main() {
	factory := NewAnimalFactory()
	factory.Register(DOG, func() Animal { return new(Dog) })
	factory.Register(CAT, func() Animal { return new(Cat) })
	factory.Register(LIZARD, func() Animal { return new(Lizard) })
	factory.Register(JAVA_PROGRAMMER, func() Animal { return new(JavaProgrammer) })

	animals := []AnimalType{CAT, DOG, LIZARD, JAVA_PROGRAMMER}
	for _, t := range animals {
		a := factory.MakeAnimal(t)
		fmt.Println(a.Speak())
	}

}
