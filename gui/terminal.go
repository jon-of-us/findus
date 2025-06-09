package gui

import (
	"findus/gui/ansi"
	"fmt"
	"os"

	term "golang.org/x/term"
)

func SetupTerminal() (restoreFunc func(), err error) {
	// Switch to alternate screen buffer
	fmt.Print(ansi.AlternateScreenBuffer)
	// Clear the alternate screen
	fmt.Print(ansi.ClearScreen)
	// Move cursor to home position
	fmt.Print(ansi.CursorHome)
	// Hide cursor
	fmt.Print(ansi.HideCursor)

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		// Attempt to restore basic settings if raw mode failed
		fmt.Print(ansi.ShowCursor)       // Show cursor
		fmt.Print(ansi.MainScreenBuffer) // Switch back to main screen buffer
		return nil, fmt.Errorf("failed to set raw mode: %w", err)
	}

	restoreFunc = func() {
		// Show cursor
		fmt.Print(ansi.ShowCursor)
		// Restore terminal state
		term.Restore(fd, oldState)
		// Switch back to main screen buffer
		fmt.Print(ansi.MainScreenBuffer)
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
