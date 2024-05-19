package components

import "findus/gui/screen"

type props struct {
	minWidth   int
	minHeight  int
	mustRender bool
}

func newEmptyProps() props {
	return props{
		minWidth:   0,
		minHeight:  0,
		mustRender: false,
	}
}
func (p *props) SetMustRender(mustRender bool) *props {
	p.mustRender = mustRender
	return p
}
func (p *props) SetMinWidth(minWidth int) *props {
	p.minWidth = minWidth
	return p
}
func (p *props) SetMinHeight(minHeight int) *props {	
	p.minHeight = minHeight
	return p
}

type Component interface {
	Render(box screen.RenderBox)
	Props() *props
}
