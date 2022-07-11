package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	sum := 0

	for {
		// Read string with space delimitter
		input, err := r.ReadString(' ')
		n := strings.TrimSpace(input)
		if n == "" {
			// empty string, no int
			continue
		}

		// Convert from string to int
		num, cErr := strconv.Atoi(n)
		if cErr != nil {
			fmt.Println(cErr)
		} else {
			sum += num
		}

		// Check if no more data input
		if err == io.EOF {
			break
		}
	}

	// Print sum
	fmt.Println("Sum:", sum)
}
