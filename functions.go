package main

import "fmt"

func main() {
	printMessage("Hello Go!")
	sayGreeting("Hello", "Sasha")
	name := "Ted"
	sayGreetingWithPointers("Hello", &name)
	// value would have changed
	// as its reference was passed to the sayGreetingWithPointers function
	// passing pointers is more memory efficient
	fmt.Println(name)
	fmt.Println("Sum :", variadicSum(1, 4, 5, 7))
	fmt.Println("Sum when pointer returned :", *variadicSumReturningPointer(1, 4, 5, 7))
	fmt.Println("Sum when using return syntactic sugar :", variadicSumWithReturnSyntacticSugar(1, 4, 5, 7))

	quotient, error := divide(5.0, 0.0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Quotient :", quotient)
	}
}

func printMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting string, name string) {
	// includes a space already
	fmt.Println(greeting, name)
}

func sayGreetingWithPointers(greeting string, name *string) {
	fmt.Println(greeting, *name)
	*name = "John"
}

func variadicSum(values ...int) int {

	// values is a slice
	fmt.Println(values)

	result := 0

	// syntax to iterate through a collection
	for _, v := range values {
		result += v
	}

	return result
}

func variadicSumReturningPointer(values ...int) *int {

	result := 0

	for _, v := range values {
		result += v
	}

	return &result
}

// syntactic sugar when it comes to returning a result
func variadicSumWithReturnSyntacticSugar(values ...int) (result int) {
	result = 0

	for _, v := range values {
		result += v
	}

	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		// avoid using panic unless the program cannot continue at all
		// panic("Zero cannot be the second argument of the divide function.")
		// instead return an error
		// returning an error is a very idiomatic way of using go
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
