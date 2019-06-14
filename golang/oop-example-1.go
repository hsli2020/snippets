package main

import (
    "fmt"
)

type Animal struct {
    Name string
    mean bool
}

type AnimalSounder interface {
    MakeNoise()
}

type Dog struct {
    Animal
    BarkStrength int
}

type Cat struct {
    Basics Animal
    MeowStrength int
}

func main() {
    myDog := &Dog{
        Animal{
           "Rover", // Name
           false,   // mean
        },
        2, // BarkStrength
    }

    myCat := &Cat{
        Basics: Animal{
            Name: "Julius",
            mean: true,
        },
        MeowStrength: 3,
    }

    MakeSomeNoise(myDog)
    MakeSomeNoise(myCat)
}

func (animal *Animal) PerformNoise(strength int, sound string) {
    if animal.mean == true {
        strength = strength * 5
    }

    for voice := 0; voice < strength; voice++ {
        fmt.Printf("%s ", sound)
    }

    fmt.Println("")
}

func (dog *Dog) MakeNoise() {
    dog.PerformNoise(dog.BarkStrength, "BARK")
}

func (cat *Cat) MakeNoise() {
    cat.Basics.PerformNoise(cat.MeowStrength, "MEOW")
}

func MakeSomeNoise(animalSounder AnimalSounder) {
    animalSounder.MakeNoise()
}