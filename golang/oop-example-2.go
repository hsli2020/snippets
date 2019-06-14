package main

import (
    "fmt"
)

type HornSounder interface {
    SoundHorn()
}

type Vehicle struct {
    List [2]HornSounder
}

type Car struct {
    Sound string
}

type Bike struct {
   Sound string
}

func main() {
    vehicle := new(Vehicle)
    vehicle.List[0] = &Car{"BEEP"}
    vehicle.List[1] = &Bike{"RING"}

    for _, hornSounder := range vehicle.List {
        hornSounder.SoundHorn()
    }
}

func (car *Car) SoundHorn() {
    fmt.Println(car.Sound)
}

func (bike *Bike) SoundHorn() {
    fmt.Println(bike.Sound)
}

func PressHorn(hornSounder HornSounder) {
    hornSounder.SoundHorn()
}