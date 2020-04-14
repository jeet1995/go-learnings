package main

import "fmt"

/*
	1. These are variable blocks.
	2. Blocks help avoid cluttering when it comes to declarations.
	3. The naming convention of variables influences their visibility in Go, for example :
		3.1 Say a variable is to be exposed, make its name begin with an upper case letter.
		3.2 When the name of the variable begins with a lower case letter, its visibility
			remains local to the scope of the package.
*/
var (
	cricketerName     = "Virat Kohli"
	odiBattingAverage = 59.97
)

var (
	length = 2.0
	height = 4.5
	width  = 3.5
)

func main() {

	var i int = 42 // Declaration
	i = 43         // Re-initialization
	j := 44        // Declaration

	fmt.Println(i)
	fmt.Println(j)
	fmt.Print(cricketerName)
}
