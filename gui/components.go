package gui

import (
	"findus/config"
	"findus/gui/ansi"
	"strings"
	"unicode/utf8"
)

type Component struct {
	Content string
	// Len is the rendered length -1 for vertical components
	Len int
}

func String(text string) Component {
	return Component{
		Content: text,
		Len:     utf8.RuneCountInString(text),
	}
}

func Shortcut(text string) Component {
	return Component{
		Content: text,
		Len:     utf8.RuneCountInString(text),
	}.Styled(config.ShortcutColor)
}

func Clipped(component Component, maxLen int) Component {
	maxLen = max(maxLen, 3)
	if component.Len > maxLen {
		return Component{
			Content: "..." + component.Content[:maxLen-3],
			Len:     maxLen,
		}
	} else {
		return component
	}
}

func SideBySide(left Component, right Component, maxLen int) Component {
	padding := maxLen - left.Len - right.Len
	content := left.Content + strings.Repeat(" ", padding) + right.Content
	return Component{
		Content: content,
		Len:     maxLen,
	}
}

func NavigationModePlaceholder() Component {
	return Component{
		Content: "Navigation Mode",
		Len:     utf8.RuneCountInString("Navigation Mode"),
	}.Styled(ansi.Gray)
}

func Lines(components ...Component) Component {
	var content strings.Builder
	for i, component := range components {
		content.WriteString(component.Content)
		if i < len(components)-1 {
			content.WriteString("\n")
		}
	}
	return Component{
		Content: content.String(),
		Len:     -1,
	}
}

func Concat(components ...Component) Component {
	var content strings.Builder
	var totalLen int

	for _, component := range components {
		content.WriteString(component.Content)
		totalLen += component.Len
	}

	return Component{
		Content: content.String(),
		Len:     totalLen,
	}
}
