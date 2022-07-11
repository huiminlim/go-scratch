// With main package, will be the script
// to launch the program
package main

import (
	"demo/demo/pkg/display"
	"demo/demo/pkg/msg"
)

func main() {
	// Type the package name directly to load each required package
	msg.Hi()
	display.Display(("Hello from display"))
	msg.Exciting("Exciting!")
}
