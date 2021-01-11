package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{}
	n, error := w.Write([]byte("Hello Go!"))
	fmt.Println(n, error)
}

// an interface describes behaviour
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cs ConsoleWriter) Write(data []byte) (int, error) {
	n, error := fmt.Println(string(data))
	return n, error
}
