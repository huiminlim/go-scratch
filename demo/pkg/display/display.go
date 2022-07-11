// Part of the display package
// and cannot have main function
package display

import "fmt"

// Capitalized function name: will be exported from the package
func Display(msg string) {
	fmt.Println(msg)
}

// Lowercase function name: internal and will not be exported
//lint:ignore U1000 for demo's sake
func hello(msg string) {
	fmt.Println("Hidden")
}
