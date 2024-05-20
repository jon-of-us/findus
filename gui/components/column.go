package components

import (
	"findus/gui/screen"
	"findus/utils"
)

type column struct {
	children []Component
	props    props
}

func NewColumn(children []Component) *column {
	props := newEmptyProps()
	mustRenderChildren := []Component{}
	for _, c := range children {
		if c.Props().mustRender {
			mustRenderChildren = append(mustRenderChildren, c)
		}
	}

	minWidth := 0
	for _, artist := range mustRenderChildren {
		minWidth = utils.MaxInt(minWidth, artist.Props().minWidth)
	}

	minHeight := 0
	for _, artist := range children {
		minHeight += artist.Props().minHeight
	}

	return &column{
		children: children,
		props:    props,
	}
}

var _ Component = &column{}

func (c *column) Props() *props {
	return &c.props
}

func (c *column) Render(box screen.RenderBox) {
	width := utils.MaxInt(box.Width, c.Props().minWidth)
	height := utils.MaxInt(box.Height, c.Props().minHeight)

	remainingHeight := height
	isChildRendered := make([]bool, len(c.children))
	heightsToRender := make([]int, len(c.children))

	// add children that must render
	for i, child := range c.children {
		if child.Props().mustRender {
			isChildRendered[i] = true
			heightsToRender[i] = child.Props().minHeight
			remainingHeight -= child.Props().minHeight
		}
	}

	// add children that fit in remaining space
	for i, child := range c.children {
		childHeight := child.Props().minHeight
		if !child.Props().mustRender && childHeight < remainingHeight {
			isChildRendered[i] = true
			heightsToRender[i] = childHeight
			remainingHeight -= childHeight
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

	for i := 0; remainingHeight > 0; i = (i + 1) % len(heightsToRender) {
		if isChildRendered[i] {
			heightsToRender[i] += 1
			remainingHeight -= 1
		}
	}

	// render all children
	heightStart := 0
	for i, child := range c.children {
		if isChildRendered[i] {
			childHeight := heightsToRender[i]
			child.Render(box.SubBox(0, heightStart, width, childHeight))
			heightStart += childHeight
		}
	}
}