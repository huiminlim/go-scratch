package main

import "fmt"

func main() {
	slice := []string{"Hello", "world", "!"}

	// i = element number
	// element is value
	for i, element := range slice {

		fmt.Print(i, " ")

		// To access each rune in the element
		for _, ch := range element {
			// Printing the rune as a character and not digit
			// using %q modifier
			fmt.Printf("%q ", ch)
		}

		fmt.Println()
	}

	a := "Hello"
	for _, ch := range a {
		fmt.Printf("%q ", ch)
	}
}
