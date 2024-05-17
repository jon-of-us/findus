package gui

import "github.com/gdamore/tcell/v2"

var style tcell.Style

func Box(x0, y0, width, height int, s tcell.Screen) {

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
