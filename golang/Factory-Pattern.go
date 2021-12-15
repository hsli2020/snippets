package main

import "fmt"

type Pet interface {
	GetName() string
	GetAge() int
	GetSound() string
}

// pet is a struct that implements Pet interface and
// would be used in any animal struct that we create.
// See `Dog` and `Cat` below
type pet struct {
	name  string
	age   int
	sound string
}

func (p *pet) GetName() string  { return p.name }
func (p *pet) GetSound() string { return p.sound }
func (p *pet) GetAge() int      { return p.age }

type Dog struct {
	pet
}

type Cat struct {
	pet
}

func GetPet(petType string) Pet {
	if petType == "dog" {
		return &Dog{
			pet{
				name:  "Chester",
				age:   2,
				sound: "bark",
			},
		}
	}
	if petType == "cat" {
		return &Cat{
			pet{
				name:  "Mr. Buttons",
				age:   3,
				sound: "meow",
			},
		}
	}
	return nil
}

func describePet(pet Pet) string {
	return fmt.Sprintf("%s is %d years old. It's sound is %s",
		pet.GetName(), pet.GetAge(), pet.GetSound())
}

func main() {
	petType := "dog"

	dog := GetPet(petType)
	petDescription := describePet(dog)

	fmt.Println(petDescription)
	fmt.Println("-------------")

	petType = "cat"
	cat := GetPet(petType)
	petDescription = describePet(cat)

	fmt.Println(petDescription)
}
