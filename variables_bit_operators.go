package main

import "fmt"

func main() {

	var a int = 5
	var b int = 9

	// OR operator
	fmt.Printf("%v \n", a|b)

	// XOR operator
	fmt.Printf("%v \n", a^b)

	// AND operator
	fmt.Printf("%v \n", a&b)

	// XNOR operator
	fmt.Printf("%v \n", a&^b)

	// bit-shift operators
	fmt.Printf("%v \n", a<<3)

	// bit-shift operators
	fmt.Printf("%v \n", a>>3)

}
