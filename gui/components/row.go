package components

import (
	"findus/utils"
	"fmt"

	"github.com/gdamore/tcell/v2"
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
	func (r *row) Render(x int, y int, width int, height int, s tcell.Screen) {

		remainingWidth := width
		widthsToRender := []int{}
		// add children that must rendersaasd
		for _, child := range r.children {
			if child.Props().mustRender {
				remainingWidth -= child.Props().minWidth
				widthsToRender = append(widthsToRender, child.Props().minWidth)
			} else {
				widthsToRender = append(widthsToRender, 0)
			}
		}
		// add children that fit in remaining space
		for remainingWidth > 0 {
			for i, child := range r.children {
				childWidth := child.Props().minWidth
				if !child.Props().mustRender && childWidth > remainingWidth {
					widthsToRender[i] = childWidth
					remainingWidth -= childWidth
				}
			}
		}
		// distribute remaining space
		for i := 0; remainingWidth > 0; i = (i - 1) % len(widthsToRender) {
			widthsToRender[i] += 1
			remainingWidth -= 1
		}
		// render all children
		widthStart := 0
		for i, child := range r.children {
			childWidth := widthsToRender[i]
			fmt.Println(childWidth)
			if childWidth == 0 {
				continue
			}
			child.Render(x + widthStart, y, childWidth, height, s)
			widthStart += childWidth
		}
	}
