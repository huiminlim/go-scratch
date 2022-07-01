//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Dice rolling function
func rollDice(numDice int, numRoll int, numSide int) int {
	// Create and initialize sum
	sum := 0

	for i := 0; i < numDice; i++ {
		for r := 0; r < numRoll; r++ {
			// rand.Intn(numSide) returns number from [0,numSide-1]
			// add 1 to adjust to [1, numSide]
			sum += (rand.Intn(numSide) + 1)
		}
	}

	// Perform checks
	if numRoll == 2 && numDice == 2 {
		fmt.Println("Snake eyes")
	} else if numRoll == 7 {
		fmt.Println("Lucky 7")
	}

	if numRoll%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	return sum
}

func main() {
	// Set seed to randomize output
	rand.Seed(time.Now().UnixNano())

	// Test the function

	// Case 1: Snake eyes and even
	fmt.Println("Rolling 2 dices of 6 sides, 2 times: ", rollDice(2, 2, 6))

	// Case 2: Lucky 7
	fmt.Println("Rolling 1 dices of 10 sides, 1 times: ", rollDice(1, 1, 10))
}
