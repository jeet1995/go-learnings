package main

import "fmt"

func main() {

	// declaring a map
	statePopulations := map[string]string{
		"Illinois":   "12 million",
		"California": "22 million",
	}

	fmt.Printf("%v, %T\n", statePopulations, statePopulations)
	fmt.Printf("%v, %T\n", statePopulations["Illinois"], statePopulations["Illinois"])

	// checking for the presence of a key
	// ok gets set to false
	_, ok := statePopulations["Georgia"]

	// deleting a key from a map
	delete(statePopulations, "Georgia")

	fmt.Printf("%v %T", ok, ok)

	// using the make function to initialize a map
	statePopulationsCopy := make(map[string]string)

	// maps are passed around by reference
	statePopulationsCopy = statePopulations

	fmt.Printf("%v \n", statePopulationsCopy)
}
