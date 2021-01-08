package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	b := &a
	var c *int = &a
	// b's value is the reference to a
	fmt.Println(a, *b, &a, b)

	// *c is a dereference operation
	// &c has the location of the pointer
	fmt.Println(a, *c, &c)

	var ms *myStruct

	// prints <nil>
	fmt.Println(ms)
	ms = new(myStruct)
	// assigning a value to the struct
	(*ms).foo = 42
	fmt.Println(ms)

	// slices and maps are copied by reference
	// arrays are copied by value

}
