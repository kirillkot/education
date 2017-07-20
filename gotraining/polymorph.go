package main

import "fmt"

type AnimalSounder interface {
	MakeNoise()
}

func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}

type Animal struct {
	Name string
	mean bool
}

type Cat struct {
	Basics       Animal
	MeowStrength int
}

type Dog struct {
	Animal
	BarkStrength int
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

func main() {
	MakeSomeNoise(
		&Dog{
			Animal: Animal{
				Name: "Rover",
				mean: true,
			},
			BarkStrength: 2,
		})

	MakeSomeNoise(
		&Cat{
			Basics: Animal{
				Name: "Cat",
				mean: false,
			},
			MeowStrength: 4,
		})
}
