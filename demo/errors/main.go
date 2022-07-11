package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

// Create function that returns error on
// operating the struct "Stuff"
func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, fmt.Errorf("no element at index %v", index)
	} else {
		return s.values[index], nil
	}
}

func main() {
	stuff := Stuff{}
	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
