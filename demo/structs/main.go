package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	HasBoarded   bool
}

// Bus struct with information of which
// passenger is at the front seat
type Bus struct {
	FrontSeat Passenger
}

func main() {
	// Creating and assign method
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	// Block creation
	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2, HasBoarded: false}
		ella = Passenger{Name: "Ella", TicketNumber: 3, HasBoarded: false}
	)
	fmt.Println(bill, ella)

	// Create then assign method
	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4

	// Update field
	casey.HasBoarded = true
	bill.HasBoarded = true

	if bill.HasBoarded {
		fmt.Println("Bill has boarded the bus")
	}

	if casey.HasBoarded {
		fmt.Println("Casey has boarded the bus")
	}

	heidi.HasBoarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat of the bus")
}
