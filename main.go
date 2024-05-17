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
	defer screen.Quit(mainScreen)
	var width, height int
	
	
	var app components.Component = components.NewOutline(
		&components.Empty{},
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
