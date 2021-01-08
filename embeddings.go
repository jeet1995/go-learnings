package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // embeddings in go aka compositions
	SpeedKPH float32
	CanFly   bool
}

func main() {

	birdA := Bird{}

	birdA.Name = "Emu"
	birdA.Origin = "Australia"
	birdA.SpeedKPH = 48
	birdA.CanFly = false

	fmt.Println(birdA)

	// another way of declaring
	birdB := Bird{
		Animal:   Animal{Name: "Giraffe", Origin: "South Africa"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(birdB.Animal)
	fmt.Println(birdB.Animal.Name)
}
