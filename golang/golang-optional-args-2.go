package main

import "fmt"

type Foobar struct {
	mutable     bool
	temperature Celsius
}

type OptionFoobar func(*Foobar) error

func NewFoobar(options ...OptionFoobar) (*Foobar, error) {
	f := &Foobar{}

	// Default values...
	f.mutable = true
	f.temperature = 37

	// Option paremeters values:
	for _, op := range options {
		err := op(f)
		if err != nil {
			return nil, err
		}
	}
	return f, nil
}

func OptionReadonlyFlag(f *Foobar) error {
	f.mutable = false
	return nil
}

func OptionTemperature(t Celsius) func(f *Foobar) error {
	return func(f *Foobar) error {
		f.temperature = t
		return nil
	}
}

var BoilingWater OptionFoobar = OptionTemperature(100)

type Celsius int

func main() {
	f1, err := NewFoobar()
	fmt.Println(f1, err)

	f2, err := NewFoobar(OptionTemperature(25))
	fmt.Println(f2, err)

	f3, err := NewFoobar(OptionTemperature(10), OptionReadonlyFlag)
	fmt.Println(f3, err)

	f4, err := NewFoobar(OptionReadonlyFlag, OptionTemperature(10))
	fmt.Println(f4, err)

	f5, err := NewFoobar(BoilingWater)
	fmt.Println(f5, err)
}
