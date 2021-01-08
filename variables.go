package main

import (
	"fmt"
	"strconv"
)

/*
	1. Unused variables causes compilation to fail
	2. Go supports implicit type inference
*/

func main() {
	var i1 int = 42

	// Also a way of declaring a variable
	i2 := 45

	fmt.Printf("%v %T\n", i1, i1)

	// Makes use of implicit type inference
	fmt.Printf("%v %T\n", i2, i2)

	i3 := strconv.Itoa(i2)

	fmt.Printf("%v %T\n", i3, i3)

	// Defining floating point nos.
	// IEEE 754 standard can be used
	// to represent floating variables
	// using scientific notation
	var floatA float32 = 5.6
	var floatB float32 = 8.2

	fmt.Printf("%v, %v\n", floatA, floatB)

	// Defining complex nos.
	var complexA complex64 = 1 + 2i

	fmt.Printf("%v, %T\n", complexA, complexA)

	// Prints the real part of the complex no.
	fmt.Printf("%v %T", real(complexA), real(complexA))

	// Prints the imaginary part of the complex no.
	fmt.Printf("%v %T", imag(complexA), imag(complexA))

	stringA := "This is my first string."
	// A string can be represented as an array of bytes
	byteArrayOfA := []byte(stringA)

	fmt.Printf("%v \n", stringA)
	fmt.Printf("%v \n", byteArrayOfA)

	var runeA rune = 'c'

	// Encoding information of the rune is store as int32
	fmt.Printf("%v, %T\n", runeA, runeA)
}
