package main

import (
	"bufio" // Added for ReadRune
	"fmt"
	"io" // Added for io.EOF
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	// Switch to alternate screen buffer
	fmt.Print("\x1b[?1049h")
	// Clear the alternate screen
	fmt.Print("\x1b[2J")
	// Move cursor to home position
	fmt.Print("\x1b[H")

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		// Ensure we switch back to main screen buffer on error before exiting
		fmt.Print("\x1b[?1049l")
		fmt.Println("Error setting raw mode:", err)
		return
	}
	defer func() {
		// Switch back to main screen buffer
		fmt.Print("\x1b[?1049l")
		term.Restore(int(os.Stdin.Fd()), oldState)
	}()

	fmt.Println("Starting interactive terminal. Press ESC to quit.")
	fmt.Println("Commands: [a-z] add, [D] duplicate, [X] clear, [Backspace] delete")

	var inputBuffer []rune

	reader := bufio.NewReader(os.Stdin) // ADDED: Create a buffered reader

	for {
		// Print current buffer, clearing the line first
		// Get terminal width to clear the line properly
		width, _, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			width = 80 // Default width if size cannot be determined
		}
		fmt.Printf("\r%s", strings.Repeat(" ", width)) // Clear line
		fmt.Printf("\r> %s", string(inputBuffer))

		// _, err = os.Stdin.Read(b[:]) // REMOVED
		char, _, err := reader.ReadRune() // CHANGED: Use ReadRune()
		if err != nil {
			if err == io.EOF { // ADDED: Handle EOF
				fmt.Println("\nInput stream closed (EOF). Exiting...")
				return
			}
			fmt.Println("\nError reading input:", err)
			break
		}

		switch {
		case char == 27: // ESC key
			fmt.Println("\nExiting...")
			return
		case char == 127 || char == 8: // Backspace (127 for many terminals, 8 for others like Windows Terminal)
			if len(inputBuffer) > 0 {
				inputBuffer = inputBuffer[:len(inputBuffer)-1]
			}
		case char == 'D':
			inputBuffer = append(inputBuffer, inputBuffer...)
		case char == 'X':
			inputBuffer = []rune{}
		case char >= 'a' && char <= 'z':
			inputBuffer = append(inputBuffer, char)
			// default:
			// Optionally handle/ignore other characters
			// fmt.Printf("\nUnknown char: %v\n", char)
		}
	}
}
