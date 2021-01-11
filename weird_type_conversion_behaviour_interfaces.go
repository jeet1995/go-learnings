package main

import "fmt"

func main() {

	// for a pointer receiver the method set
	// is all the pointer receiver methods
	// and value receiver methods
	// for a value receiver the method set is
	// purely the value receiver methods
	// it is important to remember if an implemented
	// method accepts a pointer, then implement the interface
	// as a pointer
	var wc = WriterCloser{}
	fmt.Println(wc)
}

type WriterDemo interface {
	Write([]byte) (int, error)
}

type CloserDemo interface {
	Close() error
}

type WriterCloser struct{}

func (wcDemo WriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

// will cause an error when called
//func (wcDemo *WriterCloser) Write(data []byte) (int, error)  {
//	return 0, nil
//}

func (wcDemo WriterCloser) Close() error {
	return nil
}
