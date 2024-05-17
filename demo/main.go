package main

import (
	"demo/fuzzy"
	"demo/gui"

	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
)

var input []rune
var files [][]rune
var masks [][]bool
var height, width int

func main() {
	s := initApp()

	//eventLoop
	for {
		event := s.PollEvent()

		switch ev := event.(type) { // type switch https://go.dev/tour/methods/16

		case *tcell.EventResize:
			width, height = ev.Size()
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				quitApp(s)
			case tcell.KeyBackspace:
				if last := len(input) - 1; last >= 0 {
					input = input[:last]
				}
			case tcell.KeyRune:

				input = append(input, ev.Rune())
			}

		}
		updateFiles()
		render(height, width, s)
	}
}

func updateFiles() {
	entries, err := os.ReadDir("D:/Projekte")
	//projekte

	if err != nil {
		log.Fatal(err)
	}
	f := func(file fs.DirEntry) []rune {
		return []rune(strings.ToLower(file.Name()))
	}
	files, masks = fuzzy.Find(input, mapF(f, entries))
}

func render(height, width int, s tcell.Screen) {
	s.Clear()

	gui.Box(0, 0, width, height, s)

	gui.Box(1, 1, width-2, 3, s)
	gui.Label(input, 2, 2, width-4, true, false, nil, s)

	gui.List(files, 2, 4, width-4, height-5, masks, 0, s)

	s.Sync() // due to bug in Show, blank cells don't update syle
	//	s.Show()
}

func initApp() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		panic("new sreen error")
	}

	err = s.Init()
	if err != nil {
		panic("init error")
	}

	s.Clear()
	return s
}

func quitApp(s tcell.Screen) {
	s.Fini()
	os.Exit(0)
}
