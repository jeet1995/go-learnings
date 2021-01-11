package main

import (
	"fmt"
	"time"
)

func main() {
	// go routine creates an abstraction of a thread
	go sayHello()

	// added to ensure the go routine executes
	time.Sleep(100 * time.Millisecond)

	var msg string = "Hello"
	//
	//go func() {
	//	// msg is accessible here because of closures
	//	// here a dependency is created between the
	//	// anonymous go routine and msg
	//	// the existence of such a dependency can create
	//	// an issue
	//	// *usually* the go routine will not attempt to interrupt the
	//	// main thread until the call to Sleep is encountered
	//	// the above case is akin to a race condition and is not ideal
	//	fmt.Println(msg)
	//}()

	// Hello will be printed as msg is passed
	// when the msg variable is passed as an
	// argument, it is passed by value here
	// this decouples the go routine's dependency
	// on the msg variable defined in the main
	// method through the closure property
	go func(msg string) {
		fmt.Println(msg)
	}(msg)

	msg = "Bob"

	// this sync clock cycles with real time
	// and this is a bad practice, use wait groups
	// instead
	time.Sleep(1 * time.Millisecond)

}

// this won't get printed because
// the main function exited as soon
// as the go routine was spawned
func sayHello() {
	fmt.Println("Hello")
}
