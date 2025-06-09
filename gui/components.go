package gui

import (
	"findus/config"
	"unicode/utf8"
)

// Component is an interface that UI elements should implement.
// It defines methods to get the component's string content and its rendered length.
type Component struct {
	Content string
	Lenth   int
}

func NewComponent(text string) Component {
	return Component{
		Content: text,
		Lenth:   utf8.RuneCountInString(text),
	}
}

func Shortcut(text string) Component {
	return Component{
		Content: text,
		Lenth:   utf8.RuneCountInString(text),
	}.Styled(config.ShortcutColor)
}
