package components

import (
	"github.com/gdamore/tcell/v2"
)

type props struct {
	minWidth int
	minHeight int
}

func newEmptyProps() *props {
	return &props{
		minWidth: 0,
		minHeight: 0,
	}
}

type Component interface {
	Props() *props
	Render(x int, y int, width int, height int, s tcell.Screen)
}
