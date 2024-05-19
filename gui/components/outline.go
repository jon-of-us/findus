package components

import (
	"findus/gui/screen"
	"findus/utils"
)

type outline struct {
	child Component
	props props
}

func NewOutline(child Component) *outline {
	props := *child.Props()
	props.minWidth += 2
	props.minHeight += 2
	return &outline{
		child: child,
		props: props,
	}
}

var _ Component = &outline{}

func (o *outline) Render(box screen.RenderBox) {
	width := utils.MaxInt(box.Width, o.props.minWidth)
	height := utils.MaxInt(box.Height, o.props.minHeight)
	o.child.Render(box.SubBox(1, 1, width-2, height-2))

	box.Set(0, 0, '╭')
	box.Set(width-1, 0, '╮')
	box.Set(0, height-1, '╰')
	box.Set(width-1, height-1, '╯')

	for i := 1; i < width-1; i++ {
		box.Set(i, 0, '─')
		box.Set(i, height-1, '─')
	}
	for i := 1; i < height-1; i++ {
		box.Set(width-1, i, '│')
		box.Set(0, i, '│')
	}
}
func (o *outline) Props() *props {
	return &o.props
}
