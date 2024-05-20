package components

import "findus/gui/screen"

type Line struct {
	props props
}
func NewLine() *Line {
	props := newEmptyProps()
	props.minHeight = 1
	props.mustRender = true
	return &Line{
		props: props,
	}
}

var _ Component = &Line{}
func (vl *Line) Render(box screen.RenderBox) {
	for i := 0; i < box.Width; i++ {
		box.Set(i, 0, '─')
	}
}
func (vl *Line) Props() *props {
	return &vl.props
}