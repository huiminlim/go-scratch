package main

import "fmt"

// Typically, interfaces have "-er" suffix
type Preparer interface {
	PrepareDish()
}

// Dishes that can be created
type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("Chop salad")
	fmt.Println("Add dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Println("--- Dish: ", dish, " ---")
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken wings"), Salad("Caesar salad")}
	prepareDishes(dishes)
}
