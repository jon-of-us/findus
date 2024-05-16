package main

import "github.com/gdamore/tcell/v2"

var style tcell.Style = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

func labelR(letters []rune, x, y, width int, s tcell.Screen) {

	p := min(width-3, len(letters)) //const

	for i, r := range letters[:p] {
		s.SetContent(x+i, y, r, nil, style)
	}

	if len(letters) <= width {
		for i, r := range letters[p:] {
			s.SetContent(x+p+i, y, r, nil, style)
		}

	} else {
		for i := width - 3; i < width; i++ {
			s.SetContent(x+i, y, '.', nil, style)
		}

	}
}

func labelL(letters []rune, x, y, width int, s tcell.Screen) {

	if len(letters) <= width {
		for i, r := range letters {
			s.SetContent(x+i, y, r, nil, style)
		}

	} else {
		for i := 0; i < 3; i++ {
			s.SetContent(x+i, y, '.', nil, style)
		}
		for i, r := range letters[len(letters)-width+3:] {
			s.SetContent(x+3+i, y, r, nil, style)
		}
	}
}

func list(content [][]rune, x0, y0, width, height int, s tcell.Screen) {

	for i, x := range content[:min(len(content), height)] {
		labelR(x, x0, y0+i, width, s)
	}
}

func box(x0, y0, width, height int, s tcell.Screen) {

	s.SetContent(x0, y0+height-1, '╰', nil, style)
	s.SetContent(x0+width-1, y0+height-1, '╯', nil, style)
	s.SetContent(x0, y0, '╭', nil, style)
	s.SetContent(x0+width-1, y0, '╮', nil, style)

	for i := 1; i < width-1; i++ {
		s.SetContent(x0+i, y0, '─', nil, style)
		s.SetContent(x0+i, y0+height-1, '─', nil, style)
	}
	for i := 1; i < height-1; i++ {
		s.SetContent(x0+width-1, y0+i, '│', nil, style)
		s.SetContent(x0, y0+i, '│', nil, style)
	}
}
