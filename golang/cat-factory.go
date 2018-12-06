package main

import "fmt"

type Cat interface {
    Meow()
}

type Lion struct{}

func (l Lion) Meow() {
    fmt.Println("Roar")
}

type CatFactory func() Cat

func CreateLion() Cat {
    return Lion{}
}

func main() {
    lion := CreateLion()
    lion.Meow()

    var cf CatFactory = CreateLion
    fLion := cf()
    fLion.Meow()
}
