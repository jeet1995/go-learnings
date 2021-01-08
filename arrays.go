package main

import "fmt"

func main() {

	var arrayA = []int32{5, 6, 7}
	var arrayB []int32

	fmt.Printf("%v, %T \n", arrayA, arrayA)
	fmt.Printf("%v, %T \n", arrayB, arrayB)

	// prints the length of the array
	fmt.Printf("%v, %T \n", len(arrayA), len(arrayA))
	fmt.Printf("%v, %T \n", len(arrayB), len(arrayB))

	// declaring a 2d array
	var identityMatrix = [3][3]int32{[3]int32{1, 0, 0}, [3]int32{0, 1, 0}, [3]int32{0, 0, 1}}
	fmt.Printf("%v \n", identityMatrix)

	// easier way to define an identity matrix
	var identityMatrixSimple [3][3]int32
	identityMatrixSimple[0] = [3]int32{1, 0, 0}
	identityMatrixSimple[1] = [3]int32{0, 1, 0}
	identityMatrixSimple[2] = [3]int32{0, 0, 1}

	fmt.Printf("%v \n", identityMatrixSimple)

	arrayC := [3]int32{1, 3, 5}

	// a copy by value is done
	arrayD := arrayC

	fmt.Printf("%v \n", arrayD)

	arrayE := &arrayC

	// a copy by reference is done
	fmt.Printf("%v \n", arrayE)

	arrayF := [...]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// a slice stores the reference to the underlying array
	// elements from indices 1 until but not including 3
	sliceA := arrayF[1:3]

	// elements till indices 3
	sliceB := arrayF[:3]

	// all elements
	sliceC := arrayF[:]

	fmt.Printf("%v \n", sliceA)
	fmt.Printf("%v \n", sliceB)
	fmt.Printf("%v \n", sliceC)

}
