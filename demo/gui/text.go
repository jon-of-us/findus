package gui

import "github.com/gdamore/tcell/v2"

func Label(letters []rune, x, y, width int, rightOverfolow, inverted bool, highlightMask []bool, s tcell.Screen) {
	//dots should not be highlighted
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

		var style tcell.Style
		switch {
		case highlighted && inverted:
			style = style.Background(tcell.ColorRed).Foreground(tcell.ColorReset).Reverse(true)
		case highlighted && !inverted:
			style = style.Foreground(tcell.ColorRed).Background(tcell.ColorReset).Reverse(false)
		default:
			style = style.Background(tcell.ColorReset).Foreground(tcell.ColorReset).Reverse(inverted)
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
