package main

import "fmt"

func main() {
	fmt.Println("start")

	// defer takes precedence over panic
	defer fmt.Println("this was deferred")

	panic("something bad happened")
	fmt.Println("end")
}
