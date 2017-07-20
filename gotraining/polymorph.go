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

func (dog *Dog) MakeNoise() {
	barkStrength := dog.BarkStrength

	if dog.mean {
		barkStrength = barkStrength * 5
	}

	for bark := 0; bark < barkStrength; bark++ {
		fmt.Println("BARK")
	}

	fmt.Println(" ")
}

func (cat *Cat) MakeNoise() {
	meowStrength := cat.MeowStrength

	if cat.Basics.mean == true {
		meowStrength = meowStrength * 5
	}

	for meow := 0; meow < meowStrength; meow++ {
		fmt.Printf("MEOW ")
	}

	fmt.Println("")
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
}
