package main

import "fmt"

func main() {

	sliceA := make([]int32, 5, 100) // creating a slice

	fmt.Printf("%v \n", sliceA)
	fmt.Printf("%v \n", len(sliceA)) // a slice need not have the same size throughout its lifetime
	fmt.Printf("%v \n", cap(sliceA))

	sliceA = append(sliceA, 2)

	fmt.Printf("%v \n", sliceA)
	fmt.Printf("%v \n", len(sliceA)) // a slice need not have the same size throughout its lifetime
	fmt.Printf("%v \n", cap(sliceA))

	sliceA = append(sliceA, 2, 3, 4, 5) // append is a variadic function, a function which accepts variable no. of arguments

	fmt.Printf("%v \n", sliceA)
	fmt.Printf("%v \n", len(sliceA)) // a slice need not have the same size throughout its lifetime
	fmt.Printf("%v \n", cap(sliceA))

}
