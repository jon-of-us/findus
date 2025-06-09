package gui

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func SetupTerminal() (restoreFunc func(), err error) {
	// Switch to alternate screen buffer
	fmt.Print("\\x1b[?1049h")
	// Clear the alternate screen
	fmt.Print("\\x1b[2J")
	// Move cursor to home position
	fmt.Print("\\x1b[H")
	// Hide cursor
	fmt.Print("\\033[?25l")

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		// Attempt to restore basic settings if raw mode failed
		fmt.Print("\\033[?25h")   // Show cursor
		fmt.Print("\\x1b[?1049l") // Switch back to main screen buffer
		return nil, fmt.Errorf("failed to set raw mode: %w", err)
	}

	restoreFunc = func() {
		// Show cursor
		fmt.Print("\\033[?25h")
		// Restore terminal state
		term.Restore(fd, oldState)
		// Switch back to main screen buffer
		fmt.Print("\\x1b[?1049l")
	}

	return restoreFunc, nil
}

func GetTerminalDimensions() (width, height int, err error) {
	fds := []int{int(os.Stdin.Fd()), int(os.Stdout.Fd()), int(os.Stderr.Fd())}

	for _, fd := range fds {
		width, height, err = term.GetSize(fd)
		if err == nil {
			return width, height, nil // Success
		}
	}

	// If all attempts failed, return default values and the last error
	return 80, 24, fmt.Errorf("failed to get terminal size on Stdin, Stdout, and Stderr: %w", err)
}
