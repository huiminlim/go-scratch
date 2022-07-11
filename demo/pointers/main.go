package main

import "fmt"

type Counter struct {
	hits int
}

// Function that accepts pointer
// can use '.' to access members directly (no need to dereference)
func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	// Dereference the pointer and assign new string to that
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	// Replaces the hello variable with "Hi" string
	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}     // create array by copy
	replace(&phrase[1], "Go!", &counter) // accessing copied (new) var "world" using slice
	fmt.Println(phrase)
}
