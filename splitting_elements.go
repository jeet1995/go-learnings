package main

import "fmt"

func main() {

	arrayToSplit := []int32{1, 2, 3, 4, 5}

	// spread operator has to be used
	newArray := append(arrayToSplit[:2], arrayToSplit[3:]...)

	fmt.Printf("%v \n", newArray)

	// displays unexpected results because of presence of another reference
	fmt.Printf("%v \n", arrayToSplit)
}
