package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	var msg = "Hello"

	// Add adds delta, which may be negative, to the WaitGroup counter.
	// If the counter becomes zero, all goroutines blocked on Wait are released.
	// If the counter goes negative, Add panics.
	//
	// Note that calls with a positive delta that occur when the counter is zero
	// must happen before a Wait. Calls with a negative delta, or calls with a
	// positive delta that start when the counter is greater than zero, may happen
	// at any time.
	// Typically this means the calls to Add should execute before the statement
	// creating the goroutine or other event to be waited for.
	// If a WaitGroup is reused to wait for several independent sets of events,
	// new Add calls must happen after all previous Wait calls have returned.
	waitGroup.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		// Done decrements the WaitGroup counter by one.
		waitGroup.Done()
	}(msg)
	msg = "Goodbye"

	// Wait blocks until the WaitGroup counter is zero.
	waitGroup.Wait()
}
