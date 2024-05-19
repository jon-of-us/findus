package components

import (
	"findus/gui/screen"
	"findus/utils"
)

type row struct {
	children []Component
	props    props
}
func NewRow(children []Component) *row {
	props := newEmptyProps()
		mustRenderChildren := []Component{}
			for _, c := range children {
				if c.Props().mustRender {
					mustRenderChildren = append(mustRenderChildren, c)
				}
			}

		minWidth := 0
			for _, artist := range mustRenderChildren {
				minWidth += artist.Props().minWidth
			}
		minHeight := 0
			for _, artist := range children {
				minHeight = utils.MaxInt(minHeight, artist.Props().minHeight)
			}
	return &row{
		children: children,
		props:    props,
	}
}

var _ Component = &row{}
func (r *row) Props() *props {
	return &r.props
}
func (r *row) Render(box screen.RenderBox) {
	width := utils.MaxInt(box.Width, r.Props().minWidth)
	height := utils.MaxInt(box.Height, r.Props().minHeight)

	remainingWidth := width
	isChildRendered := make([]bool, len(r.children))
	widthsToRender := make([]int, len(r.children))
	// add children that must renders
	for i, child := range r.children {
		if child.Props().mustRender {
			isChildRendered[i] = true
			widthsToRender[i] = child.Props().minWidth
			remainingWidth -= child.Props().minWidth
		} 
	}
	// add children that fit in remaining space
	for i, child := range r.children {
		childWidth := child.Props().minWidth
		if !child.Props().mustRender && childWidth < remainingWidth {
			isChildRendered[i] = true
			widthsToRender[i] = childWidth
			remainingWidth -= childWidth
		}
	}
	// distribute remaining space
	someRendered := false
		for _, rendered := range isChildRendered {
			someRendered = someRendered || rendered
		}
	if !someRendered {
		return
	}
	for i := 0; remainingWidth > 0; i = (i + 1) % len(widthsToRender) {
		if isChildRendered[i] {
			widthsToRender[i] += 1
			remainingWidth -= 1
		}
	}
	// render all children
	widthStart := 0
	for i, child := range r.children {
		if isChildRendered[i] {
			childWidth := widthsToRender[i]
			child.Render(box.SubBox(widthStart, 0, childWidth, height))
			widthStart += childWidth
		}
	}
}
