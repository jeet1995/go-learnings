package main

import "fmt"

func main() {
	// declaring a bool variable
	// note: initialized to false by default
	var boolVariable bool

	// initializing a bool variable
	boolVariable = true

	// printing a bool variable
	fmt.Printf("%T %v", boolVariable, boolVariable)

	var a uint16 = 16
	var b uint16 = 33

	//
	fmt.Print(a + b)

}
