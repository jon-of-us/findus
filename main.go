package main

import (
	"findus/gui/components"
	"findus/gui/screen"

	"github.com/gdamore/tcell/v2"
)

func main() {
	start()
}

func start() {
	mainScreen := screen.Init()
	var width, height int

	empty1 := components.NewEmpty()
	empty1.Props().SetMustRender(true).SetMinWidth(20)

	empty2 := components.NewEmpty()
	empty2.Props().SetMinWidth(25)

	empty3 := components.NewEmpty()
	empty3.Props().SetMinWidth(30).SetMustRender(true)
	
	var app components.Component = components.NewRow(
		[]components.Component{
			components.NewOutline(
				empty1,
			),
			components.NewOutline(
				empty2,
			),
			components.NewOutline(
				empty3,
			),
		},
	)
	//eventLoop
	for {
		event := mainScreen.PollEvent()
		switch ev := event.(type) { // type switch https://go.dev/tour/methods/16
			case *tcell.EventResize:
				width, height = ev.Size()
			case *tcell.EventKey:
				switch ev.Key() {
					case tcell.KeyEscape:
						screen.Quit(mainScreen)
				}
		}
		mainScreen.Fill(' ', tcell.StyleDefault)
		mainScreen.Clear()
		app.Render(screen.NewRenderBox(mainScreen, 0, 0, width, height))
		mainScreen.Show()
	}
}
