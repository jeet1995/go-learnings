package main

import "fmt"

func main() {
	fmt.Println("start")
	// execution pushed beyond that of the
	// last statement just before the return
	// if all statements are deferred, then they
	// execute in LIFO order
	defer fmt.Println("middle")
	fmt.Println("end")

	varA := "start"
	// will print "start"
	// it takes the value at the position it
	// would have been executed at without the
	// defer keyword
	// 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
	// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
	// 3. Deferred functions may read and assign to the returning function's named return values.
	defer fmt.Printf(varA)
	varA = "end"

}
