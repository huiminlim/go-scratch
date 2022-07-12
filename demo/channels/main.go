package main

import (
	"fmt"
	"time"
)

type ControlMessage int

type Job struct {
	data   int
	result int
}

const (
	DoExit = iota
	ExitOk
)

// Takes in 2 channels of type Job
func doubler(jobs, results chan Job, control chan ControlMessage) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			default:
				panic("Unhandled control message")
			}
		case job := <-jobs:
			results <- Job{data: job.data, result: job.data * 2}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	// Create channels
	jobs := make(chan Job, 50)
	results := make(chan Job, 50)
	control := make(chan ControlMessage)

	// Launch goroutine
	go doubler(jobs, results, control)

	// Create jobs and push to channel
	for i := 0; i < 30; i++ {
		jobs <- Job{i, 0}
	}

	// Get return values from channel
	for {
		select {

		// If result value returned, print out
		case result := <-results:
			fmt.Println(result)

		// Give goroutine 500ms to compute
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")

			// Exit goroutine
			control <- DoExit
			<-control

			fmt.Println("program exit")
			return
		}
	}
}
