package components

import "github.com/gdamore/tcell/v2"

type Empty struct{}

func (*Empty) Render(x int, y int, width int, height int, s tcell.Screen) {}
func (*Empty) Props() *props{ return newEmptyProps() }
