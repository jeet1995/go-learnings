package main

import "fmt"

func main() {
	// overlapping cases cause syntax errors
	switch 1 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}

	// an expression can be checked against cases
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 3, 4:
		fmt.Println("two, three or four")
	default:
		fmt.Println("another number")
	}

	i := 10

	// overlapping is allowed here
	// fallthrough executes the next case as well
	switch {
	case i <= 10:
		fmt.Println("The value is less than 10.")
		fallthrough
	case i <= 20:
		fmt.Println("The value is less than 20.")
	default:
		fmt.Println("The value is not less than 20")
	}

	// type switch
	// for arrays to be of the same type
	// the type of elements they use and
	// their size have to be the same
	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("i is an int")
		break
		// unreachable code
		fmt.Println("Will not be executed as it follows a break")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}

}
