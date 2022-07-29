package main

import (
	"fmt"
)

type Options struct {
	Name string
	Age  int
}

type OptionFunc func(*Options)

func WithName(name string) OptionFunc {
	return func(o *Options) {
		o.Name = name
	}
}

func WithAge(age int) OptionFunc {
	return func(o *Options) {
		o.Age = age
	}
}

func Greet(options ...OptionFunc) string {
	opts := Options{
		Name: "Aiden",
		Age:  30,
	}
	for _, o := range options {
		o(&opts)
	}
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", opts.Name, opts.Age)
}

func main() {
	var greeting string

	greeting = Greet() // Greet(Name="Aiden", Age=30)
	fmt.Println(greeting)

	greeting = Greet(WithName("Alice"), WithAge(20))
	fmt.Println(greeting)
}
