package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	fmt.Printf("Running computation with %v and %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	// Passing function as function literal
	fmt.Println("2 + 2 = ", compute(2, 2, add))

	// Inline function literal
	fmt.Println("10 - 2 = ", compute(10, 2, func(lhs, rhs int) int { return lhs - rhs }))

	// Passing closure as function literal
	mul := func(lhs, rhs int) int {
		fmt.Printf("Multiply %v and %v = ", lhs, rhs)
		return rhs * lhs
	}
	fmt.Println(compute(3, 3, mul))
}
