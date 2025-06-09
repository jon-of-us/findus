package ansi

// Style represents a text style using ANSI escape sequences.
type Style string

// Style constants for colors, bold, and terminal control sequences.
const (
	Reset    string = "\x1b[0m"
	Bold     Style  = "\x1b[1m"
	Black    Style  = "\x1b[30m"
	Red      Style  = "\x1b[31m"
	Green    Style  = "\x1b[32m"
	Yellow   Style  = "\x1b[33m"
	Blue     Style  = "\x1b[34m"
	Magenta  Style  = "\x1b[35m"
	Cyan     Style  = "\x1b[36m"
	White    Style  = "\x1b[37m"
	Gray     Style  = "\x1b[90m"
	DarkBlue Style  = "\x1b[94m"

	// Terminal control sequences
	AlternateScreenBuffer string = "\x1b[?1049h"
	MainScreenBuffer      string = "\x1b[?1049l"
	ClearScreen           string = "\x1b[2J"
	CursorHome            string = "\x1b[H"
	HideCursor            string = "\x1b[?25l"
	ShowCursor            string = "\x1b[?25h"
)
