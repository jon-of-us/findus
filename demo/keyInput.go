package main

import (
	"unicode"

	"github.com/gdamore/tcell/v2"
)

func handleKeyInput(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyEscape:
		quitApp(s)
	case tcell.KeyBackspace:
		if last := len(input) - 1; last >= 0 {
			input = input[:last]
		}
		updateSearch()
		resetScroll()
	case tcell.KeyRune:
		switch {
		case unicode.IsLower(event.Rune()):
			input = append(input, event.Rune())
			updateSearch()
			resetScroll()
		case unicode.IsUpper(event.Rune()):
			handleComandKey(event.Rune())
		}
	}
}

func handleComandKey(char rune) {
	switch char {
	case 'J':
		downEv()
	case 'K':
		upEv()
	}

}
