package main

import (
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
)

var input []rune
var width, height int
var files [][]rune

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
		render(s)
	}
}

func updateFiles() {
	entries, err := os.ReadDir("D:/Projekte")

	if err != nil {
		log.Fatal(err)
	}
	f := func(file fs.DirEntry) []rune {
		return []rune(strings.ToLower(file.Name()))
	}
	files, _ = fzFind(input, mapF(f, entries))
}

func render(s tcell.Screen) {
	w := 30
	s.Clear()

	box(0, 0, width, height, s)

	box(1, 1, w+2, 3, s)
	labelL(input, 2, 2, w, s)

	list(files, 2, 4, w, height-5, s)

	s.Show()
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
