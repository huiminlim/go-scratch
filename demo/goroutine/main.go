package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	// Closure to perform capitalization action
	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	// Run goroutine on the closure
	for i := 0; i < len(data); i++ {
		go capIt(data[i])
	}
	time.Sleep(1000 * time.Millisecond) // Wait for goroutines to complete

	fmt.Printf("CapitalizedL %c\n", capitalized)
}
