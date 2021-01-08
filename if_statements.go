package main

import "fmt"

func main() {

	if true {
		fmt.Println("The test is true.")
	}

	statePopulations := map[string]int{
		"California":   100000,
		"Texas":        5678,
		"Florida":      6789,
		"New York":     8421,
		"Pennsylvania": 1281,
		"Illinois":     2678,
		"Ohio":         1167,
	}

	// initializer syntax
	if population, ok := statePopulations["California"]; ok {
		fmt.Printf("The population of %v is : %v", "California", population)
	}

}
