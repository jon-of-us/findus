package main

import (
	"bufio"
	"findus/backend"
	"findus/gui"
	"findus/gui/ansi"
	"os"
)

type Mode int

const (
	ModeCommand Mode = iota
	ModeNavigation
)

type App struct {
	input      string
	backend    backend.State
	mode       Mode
	readWriter bufio.ReadWriter
	gui        gui.Component
	onExit     func()
}

func NewApp() *App {
	restoreTerminal := gui.SetupTerminal()
	app := &App{
		input:      "",
		backend:    *backend.NewState(),
		mode:       ModeCommand,
		readWriter: *bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)),
		onExit:     restoreTerminal,
	}

	return app
}

func (app *App) render() {
	guiComponent := gui.String("Error")
	width, _ := gui.GetTerminalDimensions()
	if app.mode == ModeNavigation {
		gui.Lines(
			gui.String("Welcome to Findus!").Bold(),
			gui.SideBySide(
				gui.String("Navigation Mode").Styled(ansi.Gray),
				gui.Concat(
					gui.String("Press "), gui.Shortcut("M"), gui.String(" to change mode"),
				),
				width,
			),
			gui.String("Input: "+app.input),
		)
	} else if app.mode == ModeCommand {
		gui.Lines(
			gui.String("Welcome to Findus!").Bold(),
			gui.SideBySide(
				gui.String("Command Mode").Styled(ansi.Gray),
				gui.Concat(
					gui.String("Press "), gui.Shortcut("M"), gui.String(" to change mode"),
				),
				width,
			),
			gui.String("Input: "+app.input),
		)

	}
	println(ansi.ClearScreen)
	println(ansi.CursorHome)
	println(guiComponent.Content)
}

func (app *App) exit() {
	app.onExit()
	app.backend.Exit()
}

// Returns if it will continue running
func (app *App) Update() bool {
	rune, _, err := app.readWriter.ReadRune()
	if err != nil {
		// Handle error, potentially EOF, so exit
		app.exit()
		return false
	}

	if (rune >= 'a' && rune <= 'z') || rune == ' ' {
		app.input += string(rune)
	} else if rune == 27 {
		app.exit()
		return false
	} else if rune == 'M' {
		if app.mode == ModeCommand {
			app.mode = ModeNavigation
		} else {
			app.mode = ModeCommand
		}
		app.input = ""
	}
	app.render()
	return true
}
