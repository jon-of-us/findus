package gui

import "fmt"

// ANSI escape code for resetting all attributes
const ansiReset = "\\x1b[0m"

// Bold wraps the input string with ANSI escape codes to make it bold.
func Bold(text string) string {
	const ansiBold = "\\x1b[1m"
	return fmt.Sprintf("%s%s%s", ansiBold, text, ansiReset)
}

// Color represents a predefined ANSI color name.
// For more advanced color handling, one might use direct ANSI codes or a more complex color type.
type Color string

const (
	ColorBlack   Color = "black"
	ColorRed     Color = "red"
	ColorGreen   Color = "green"
	ColorYellow  Color = "yellow"
	ColorBlue    Color = "blue"
	ColorMagenta Color = "magenta"
	ColorCyan    Color = "cyan"
	ColorWhite   Color = "white"
)

var colorToAnsi = map[Color]string{
	ColorBlack:   "\\x1b[30m",
	ColorRed:     "\\x1b[31m",
	ColorGreen:   "\\x1b[32m",
	ColorYellow:  "\\x1b[33m",
	ColorBlue:    "\\x1b[34m",
	ColorMagenta: "\\x1b[35m",
	ColorCyan:    "\\x1b[36m",
	ColorWhite:   "\\x1b[37m",
}

// Colored wraps the input string with ANSI escape codes for the specified foreground color.
// If the colorName is not a predefined color, the original string is returned without coloring.
func Colored(text string, colorName Color) string {
	ansiColorCode, ok := colorToAnsi[colorName]
	if !ok {
		// If the color name is not found, return the original text without color
		// Alternatively, one could return an error or use a default color.
		return text
	}
	return fmt.Sprintf("%s%s%s", ansiColorCode, text, ansiReset)
}
