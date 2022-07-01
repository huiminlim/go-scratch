// Using main package
package main

// Import Format package to print information
import "fmt"

func main() {
	// Assigning string variable and printing out using fmt package
	var firstName = "Bob"       // type inference
	var lastName string = "Jon" // type annotation
	fmt.Println("Hello! My name is ", firstName, lastName)

	// Create and assign
	greeting := "Hello, world!"
	fmt.Println(greeting)

	// Uninitialized variable
	var sum int
	fmt.Println("sum is ", sum)

	// Comma ok syntax
	part1, other := 1, 4
	fmt.Println("part1: ", part1, ", other: ", other)
	part2, other := 12, 11
	fmt.Println("part1: ", part2, ", other: ", other)

	// Block declaration
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println(lessonName, lessonType)
}
