package components

import (
	"github.com/gdamore/tcell/v2"
)

type outline struct {
	child Component
}
func NewOutline(child Component) outline {
	return outline{
		child: child,
	}
}

func (o outline) Render(x int, y int, width int, height int, s tcell.Screen) {
	o.child.Render(x - 1, y - 1, width - 2, height - 2, s)
	
	s.SetContent(x, y+height-1, '╰', nil, style)
	s.SetContent(x+width-1, y+height-1, '╯', nil, style)
	s.SetContent(x, y, '╭', nil, style)
	s.SetContent(x+width-1, y, '╮', nil, style)

	for i := 1; i < width-1; i++ {
		s.SetContent(x+i, y, '─', nil, style)
		s.SetContent(x+i, y+height-1, '─', nil, style)
	}
	for i := 1; i < height-1; i++ {
		s.SetContent(x+width-1, y+i, '│', nil, style)
		s.SetContent(x, y+i, '│', nil, style)
	}
}
func (o outline) Props() *props {
	return &props{
		o.child.Props().minHeight + 2,
		o.child.Props().minHeight + 2,
	}
}

