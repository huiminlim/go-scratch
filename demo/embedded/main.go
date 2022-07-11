package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

// Creating interface #1
type Conveyor interface {
	Convey() BeltSize
}

// Creating interface #2
type Shipper interface {
	Ship() Shipping
}

// Creating embedded interface
type WarehouseAutomator interface {
	Conveyor
	Shipper
}

// Create a type that implement embedded interface
type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}
func (s *SpamMail) Ship() Shipping {
	return Air
}
func (s *SpamMail) Convey() BeltSize {
	return Small
}

// Create function with embedded type
func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

// Create type that is not embedded type
type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := SpamMail{4000}
	automate(&mail)

	waste := ToxicWaste{300}
	fmt.Println("Waste weight:", waste.weight)
	// automate(&waste) // error, can't use because not implemented
}
