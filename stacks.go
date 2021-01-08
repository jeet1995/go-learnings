package main

func main() {
	stack := []int32{1, 2, 3, 4}

	// pushing an element
	stack = append(stack, 3)

	// popping an element
	stack = stack[:len(stack)-1]
}
