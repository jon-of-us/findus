package gui

import (
	"findus/gui/ansi"
	"fmt"
	"os"

	term "golang.org/x/term"
)

func SetupTerminal() (restoreFunc func()) {
	fmt.Print(ansi.AlternateScreenBuffer)
	fmt.Print(ansi.ClearScreen)
	fmt.Print(ansi.CursorHome)
	fmt.Print(ansi.HideCursor)

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		// Attempt to restore basic settings if raw mode failed
		fmt.Print(ansi.ShowCursor)
		fmt.Print(ansi.MainScreenBuffer)
		panic(fmt.Errorf("failed to set raw mode: %w", err))
	}

	restoreFunc = func() {
		fmt.Print(ansi.ShowCursor)
		term.Restore(fd, oldState)
		fmt.Print(ansi.MainScreenBuffer)
	}

	return restoreFunc
}

func GetTerminalDimensions() (width, height int) {
	fds := []int{int(os.Stdin.Fd()), int(os.Stdout.Fd()), int(os.Stderr.Fd())}

	for _, fd := range fds {
		width, height, err := term.GetSize(fd)
		if err == nil {
			return width, height
		}
	}

	return 40, 10 // Default fallback dimensions
}
