package components

import "github.com/gdamore/tcell/v2"

type empty struct{
	props props
}
func NewEmpty() *empty {
	return &empty{
		props: newEmptyProps(),
	}
}

var _ Component = &empty{}
func (*empty) Render(x int, y int, width int, height int, s tcell.Screen) {}
func (e *empty) Props() *props{ 
	return &e.props
}
