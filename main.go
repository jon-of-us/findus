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
	
	var app components.Component = components.NewRow(
		[]components.Component{
			components.NewOutline(
				components.NewEmpty(),
			),
			components.NewOutline(
				components.NewEmpty(),
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
		mainScreen.Clear()
		app.Render(0,0,width,height,mainScreen)
		mainScreen.Show()
	}
}
