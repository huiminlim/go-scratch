package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		fmt.Print(room.name, " ")
		if room.cleaned {
			fmt.Println("is cleaned")
		} else {
			fmt.Println("is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms)

	fmt.Println("Performing cleaning...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true

	checkCleanliness(rooms)
}
