package main

import "fmt"

type Parent struct{}

func (p Parent) DoThing() {
    fmt.Println("Done!")
}

type Child struct {
	Parent // unnamed struct member
}

func main() {
	c := Child{}
	c.DoThing()
}
