package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	(&g).greetMethodWithPointer()
	fmt.Println(g.name)
}

type greeter struct {
	greeting string
	name     string
}

// this function is a method
// as it specifies an additional context
// a context is any type the method can
// execute on
// h is the value receiver
func (h greeter) greet() {
	fmt.Println(h.greeting, h.name)
}

// this function works with the pointer
// context type
func (h *greeter) greetMethodWithPointer() {
	fmt.Println(h.name, h.greeting)
	h.name = "Foden"
}
