package screen

import (
	"findus/utils"

	"github.com/gdamore/tcell/v2"
)

// RenderBox is a struct that represents a box on the screen that can be rendered to.
// it prevents components from rendering outside of their bounds
type RenderBox struct {
	screen tcell.Screen
	X int
	Y int
	Width int
	Height int
	defaultStyle tcell.Style
}
func NewRenderBox(screen tcell.Screen, x int, y int, width int, height int) RenderBox {
	return RenderBox{
		screen: screen,
		X: x,
		Y: y,
		Width: width,
		Height: height,
		defaultStyle: tcell.StyleDefault.
			Background(tcell.ColorReset).
			Foreground(tcell.ColorReset),
	}
}
func (r *RenderBox) SubBox(x int, y int, width int, height int) RenderBox {
	newX, newWidth := overlap(r.X, r.Width, r.X + x, width)
	newY, newHeight := overlap(r.Y, r.Height, r.Y + y, height)
	return NewRenderBox(r.screen, newX, newY, newWidth, newHeight)
}

func (r *RenderBox) SetStyle(style tcell.Style) {
	r.defaultStyle = style
}
func (r *RenderBox) Set(x int, y int, runeToSet rune) {
	if x < 0 || x >= r.Width || y < 0 || y >= r.Height {
		return
	}
	r.screen.SetContent(r.X + x, r.Y + y, runeToSet, nil, r.defaultStyle)
}

func overlap(start1, width1, start2, width2 int) (int, int) {
	newStart := utils.MaxInt(start1, start2)
	newEnd := utils.MinInt(start1 + width1, start2 + width2)
	newWidth := utils.MaxInt(0, newEnd - newStart)
	return newStart, newWidth
}