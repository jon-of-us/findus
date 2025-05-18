package gui

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type Terminal struct {
	oldState     *term.State
	currentLines int
}

func NewTerminal() (*Terminal, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	fmt.Print("\033[?25l") // Hide cursor
	return &Terminal{
		oldState:     oldState,
		currentLines: 0,
	}, nil
}

func (t *Terminal) Write(s fmt.Stringer) {
	str := s.String()
	fmt.Print(str)
	t.currentLines += strings.Count(str, "\n")
}

func (t *Terminal) Clear() {
	if t.currentLines > 0 {
		fmt.Printf("\033[%dA", t.currentLines)
	}
	fmt.Print("\033[J")
	t.currentLines = 0
}

func (t *Terminal) Restore() {
	t.Clear()
	fmt.Print("\033[?25h") // Show cursor
	if t.oldState != nil {
		term.Restore(int(os.Stdin.Fd()), t.oldState)
	}
}

func GetTerminalDimensions() (width, height int, err error) {
	fd := int(os.Stdin.Fd())
	width, height, err = term.GetSize(fd)
	if err != nil {
		fd = int(os.Stdout.Fd())
		width, height, err = term.GetSize(fd)
		if err != nil {
			fd = int(os.Stderr.Fd())
			width, height, err = term.GetSize(fd)
		}
	}

	if err != nil {
		return 80, 24, err
	}
	return width, height, nil
}
