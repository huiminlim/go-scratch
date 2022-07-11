package main

import "fmt"

// Variadic functions accept unknown number of input parameters
// here, "nums" are slices
func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	// Variadics with multiple int input parameters
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// Variadics with 2 slices of arrays
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	merge := append(a, b...)   // merging 2 slices
	fmt.Println(sum(merge...)) // expanding merged slice with ...
}
