package components

import (
	"github.com/gdamore/tcell/v2"
)

type props struct {
	minWidth int
	minHeight int
	mustRender bool	
}
func newEmptyProps() props {
	return props{
		minWidth: 0,
		minHeight: 0,
		mustRender: false,
	}
}
func (p *props) SetMustRender(mustRender bool) {
	p.mustRender = mustRender
}

type Component interface {
	Render(x int, y int, width int, height int, s tcell.Screen)
	Props() *props
}

