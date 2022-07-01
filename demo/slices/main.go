package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
	fmt.Println()
}

func main() {
	// Create first route
	route := []string{"Grocery", "Department Store", "Salon"}
	printSlice("Route 1", route)

	// Append another location to slice
	route = append(route, "Home")
	printSlice("Route 2", route)

	// Visit certain locations in the route
	route = route[2:]
	printSlice("Selected locations", route)
}
