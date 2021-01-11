package main

import "fmt"

func main() {

	myInt := IntCounter(0)

	// Increment works in the context
	// of the pointer of type IntCounter
	var inc Incrementer = &myInt

	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
