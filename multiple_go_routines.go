package main

import (
	"fmt"
	"sync"
)

var waitGroupNew = sync.WaitGroup{}
var counter = 0

// here the counter is not printed in order
// we need a mutex
func main() {
	for i := 0; i < 10; i++ {
		waitGroupNew.Add(2)
		// spawn thread
		go printCounter()

		// spawn thread
		go increment()
	}
	// placed here so that the main
	// function does not exit from here
	// to early
	waitGroupNew.Wait()
}

func printCounter() {
	fmt.Printf("Hello #%v\n", counter)
	waitGroupNew.Done()
}

func increment() {
	counter++
	waitGroupNew.Done()
}
