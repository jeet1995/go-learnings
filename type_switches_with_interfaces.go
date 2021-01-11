package main

import "fmt"

func main() {

	var i interface{} = true

	switch i.(type) {

	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("the type of i could not be determined")
	}

}
