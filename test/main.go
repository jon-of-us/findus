package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {

	s, err := tcell.NewScreen()
	if err != nil {
		panic("new sreen error")
	}

	err = s.Init()
	if err != nil {
		panic("init error")
	}

	s.PollEvent()
	s.PollEvent()

	var style tcell.Style

	style = style.Foreground(tcell.ColorRed).Background(tcell.ColorDefault).Reverse(false)
	s.SetContent(0, 0, 'x', nil, style)
	s.Show()

	s.PollEvent()

	style = style.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault).Reverse(true)
	s.SetContent(0, 0, ' ', nil, style)
	s.Show()

	s.PollEvent()

	s.Fini()
	os.Exit(0)
}

//style = style.Background(tcell.ColorRed).Foreground(tcell.ColorReset).Reverse(true)
//case highlighted && !inverted && char != '.':
//style = style.Foreground(tcell.ColorRed).Background(tcell.ColorDefault).Reverse(false)
//default:
//style = style.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault).Reverse(inverted)
