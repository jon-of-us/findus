package components

import (
	"findus/gui/screen"
)

type empty struct{
	props props
}
func NewEmpty() *empty {
	return &empty{
		props: newEmptyProps(),
	}
}

var _ Component = &empty{}
func (*empty) Render(box screen.RenderBox ) {
	// set all to #
	for x := 0; x < box.Width; x++ {
		for y := 0; y < box.Height; y++ {
			box.Set(x, y, '▒')
		}
	}
}
func (e *empty) Props() *props{ 
	return &e.props
}
