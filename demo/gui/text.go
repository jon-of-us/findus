package gui

import "github.com/gdamore/tcell/v2"

func Label(letters []rune, x, y, width int, rightOverfolow, inverted bool, highlightMask []bool, s tcell.Screen) {
	for i := 0; i < width; i++ {

		var char rune
		switch {
		case len(letters) > width && ((!rightOverfolow && i < 3) || (rightOverfolow && i >= width-3)):
			char = '.'
		case i < len(letters):
			char = letters[i]
		default:
			char = ' '
		}

		highlighted := i < len(highlightMask) && highlightMask[i]

		var style tcell.Style = tcell.StyleDefault
		switch {
		case highlighted && inverted && char != '.':
			style = style.Background(tcell.ColorRed).Foreground(tcell.ColorReset).Reverse(true)
		case highlighted && !inverted && char != '.':
			style = style.Foreground(tcell.ColorRed).Background(tcell.ColorDefault).Reverse(false)
		default:
			style = style.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault).Reverse(inverted)
		}
		s.SetContent(x+i, y, char, nil, style)

	}
}

func List(lines [][]rune, x, y, width, height int, boldMasks [][]bool, selected int, s tcell.Screen) {

	for i, line := range lines[:min(len(lines), height)] {
		inverted := i == selected
		var boldMask []bool
		if len(boldMasks) > i {
			boldMask = boldMasks[i]
		} else {
			boldMask = nil
		}
		Label(line, x, y+i, width, true, inverted, boldMask, s)
	}
}

func Scroll(length, height int) (func(), func(), func() (int, int, int)) {
	pos, scroll := 0, 0
	height = min(length, height)

	up := func() {
		switch {
		case pos == scroll && pos != 0:
			pos--
			scroll--
		case pos > scroll:
			pos--
		}
	}
	down := func() {
		switch {
		case pos == scroll+height-1 && pos != length-1:
			pos++
			scroll++
		case pos < scroll+height-1:
			pos++
		}
	}
	get := func() (int, int, int) {
		return scroll, scroll + height, pos - scroll
	}
	return up, down, get
}
