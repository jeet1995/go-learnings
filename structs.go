package main

import "fmt"

// casing rules to imply whether
// a variable is to be exported
// across a package apply here as
// well
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {

	// positional syntax can be used as well
	// but do not use this as it can cause
	// a maintenance problem
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Printf("%v, %T\n", aDoctor, aDoctor)

	// an anonymous struct
	bDoctor := struct {
		name string
	}{name: "John Pertwee"}

	fmt.Printf("%v, %T\n", bDoctor, bDoctor)

	// structs are copied by value
	cDoctor := bDoctor
	cDoctor.name = "Jon Lewis"
	// bDoctor will not get modified
	fmt.Printf("%v, %T\n", bDoctor, bDoctor)
	fmt.Printf("%v, %T\n", cDoctor, cDoctor)

}
