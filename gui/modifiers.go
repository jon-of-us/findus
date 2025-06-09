package gui

import (
	"findus/gui/ansi"
	"fmt"
)

// Bold applies the bold style to the component.
func (c Component) Bold() Component {
	return Component{
		Content: fmt.Sprintf("%s%s%s", ansi.Bold, c.Content, ansi.Reset),
		Lenth:   c.Lenth,
	}
}

// Styled applies a given style to the component.
func (c Component) Styled(style ansi.Style) Component {
	return Component{
		Content: fmt.Sprintf("%s%s%s", style, c.Content, ansi.Reset),
		Lenth:   c.Lenth,
	}
}
