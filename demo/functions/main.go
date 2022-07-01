package main

import "fmt"

// Simple sum function
func sum(a int, b int) int {
	return a + b
}

// Multiple return function
func many() (int, int, int) {
	return 1, 2, 3
}

func main() {
	// Testing simple sum function
	total := sum(10, 2)
	fmt.Println("10 + 2 = ", total)

	// Testing multiple return function
	a, b, c := many()
	fmt.Println(a, b, c)
}
