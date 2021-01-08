package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]

	// go does not allow pointer arithmetic
	// d := &a[1] - 8

	fmt.Printf("%v %p %p \n", a, b, c)

}
