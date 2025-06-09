package gui

import "fmt"

// Render clears the terminal screen and then prints the provided content.
func Render(content string) {
	// Clear the entire screen
	fmt.Print("\\x1b[2J")
	// Move cursor to the home position (top-left)
	fmt.Print("\\x1b[H")
	// Print the content
	fmt.Print(content)
}
