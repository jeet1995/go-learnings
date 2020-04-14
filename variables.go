package main

import "fmt"

/*
	1. Unused variables causes compilation to fail
	2. Go supports implicit type inference
*/

func main()  {


	var i1 int = 42
	i2 := 45

	fmt.Printf("%v %T\n", i1, i1)

	// Makes use of implicit type inference
	fmt.Printf("%v %T\n", i2, i2)

}
