package main

import "fmt"

func main() {

	// func as an object
	var f func() = func() {
		fmt.Println("Hello Go!")
	}
	f()

	var divide func(float64, float64) (float64, error)

	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}

		return a / b, nil
	}

	d, err := divide(5.0, 3.0)

	fmt.Println(d, err)
}
