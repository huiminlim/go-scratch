package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	counter := 0

	// Create wait group
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		counter += 1

		// Add one to wait group
		wg.Add(1)

		go func() {
			defer func() {
				fmt.Println(counter, "goroutine remaining")
				counter -= 1

				// Indicate that goroutine is complete
				wg.Done()
			}()
			duration := time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Println("Waiting for", duration)
			time.Sleep(duration)
		}()
	}
	fmt.Println("Waiting for goroutine to finish")

	// blocks until all counters in wait group are 0
	wg.Wait()

	fmt.Println("done!")
}
