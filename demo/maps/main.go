package main

import "fmt"

func main() {
	// Create map
	shoppingList := make(map[string]int)

	// Add items to shopping list map
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1 // assign zero as default, then add 1
	shoppingList["eggs"] += 3

	// Print list
	fmt.Println(shoppingList)

	// Delete key
	delete(shoppingList, "milk")
	fmt.Println("Milk deleted:", shoppingList)

	// Access key-pair
	fmt.Println("Need", shoppingList["eggs"], "eggs")

	// Check existence
	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Don't need cereal")
	} else {
		fmt.Println("Need", cereal, "cereal")
	}

	// Iterating
	sum := 0
	for _, count := range shoppingList {
		sum += count
	}
	fmt.Println("Need", sum, "items")
}
