package main

import "fmt"

func main() {
	for i := 0; i < 6; i += 2 {
		fmt.Println(i)
	}

	// with two variables
	// incrementing two variables side by side
	// is a compilation error (when using an increment
	// operator)
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// go's syntactic sugar for a while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// go's syntactic sugar for a do-while loop
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	// an alternate syntax
	i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// go's version of goto
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// iterating through collections
	arrayA := []int{1, 2, 3}
	fmt.Println(arrayA)

	for key, value := range arrayA {
		fmt.Println(key, value)
	}

	stringA := "hello go!"

	for key, value := range stringA {
		fmt.Println(key, string(value))
		fmt.Println(key, value)
	}

	statePopulations := map[string]int{
		"California":   100000,
		"Texas":        5678,
		"Florida":      6789,
		"New York":     8421,
		"Pennsylvania": 1281,
		"Illinois":     2678,
		"Ohio":         1167,
	}

	for key, value := range statePopulations {
		fmt.Println(key, value)
	}

	// _ is a write only variable
	for _, value := range statePopulations {
		fmt.Println(value)
	}
}
