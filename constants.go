package main

import (
	// "math"
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	unassigned = iota
	valA
	valB
	valC
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// Values of const types have to be
	// determined at compile-time
	// const constA float64 = math.Sin(1.57)

	const constA int32 = 9

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", a)

	fmt.Printf("%v \n", unassigned)
	fmt.Printf("%v \n", valA)
	fmt.Printf("%v \n", valB)
	fmt.Printf("%v \n", valC)

	fmt.Printf("%v \n", KB)
	fmt.Printf("%v \n", MB)
	//fmt.Printf("%v \n", GB)
	//fmt.Printf("%v \n", TB)
	//fmt.Printf("%v \n", PB)
	//fmt.Printf("%v \n", EB)
	//fmt.Printf("%v \n", ZB)

}
