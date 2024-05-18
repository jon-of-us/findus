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

// state
var s tcell.Screen
var input []rune
var files [][]rune
var masks [][]bool
var height, width int
var upEv, downEv func()
var get func() (int, int, int)

func main() {
	s = initApp()
	defer quitApp(s)

	updateSearch()

	//eventLoop
	for {
		event := s.PollEvent()

		switch ev := event.(type) { // type switch https://go.dev/tour/methods/16
		case *tcell.EventResize:
			width, height = ev.Size()
			resetScroll() //finetune needed

		case *tcell.EventKey:
			handleKeyInput(ev)
		}
		render()
	}
}

func updateSearch() {
	entries, err := os.ReadDir("D:/Projekte")

	if err != nil {
		log.Fatal(err)
	}
	f := func(file fs.DirEntry) []rune {
		return []rune(strings.ToLower(file.Name()))
	}
	files, masks = fuzzy.Find(input, mapF(f, entries))
}

func resetScroll() {
	upEv, downEv, get = gui.Scroll(len(files), max(0, height-5))
}

func render() {
	s.Clear()

	if height >= 3 {
		gui.Box(0, 0, width, height, s)
	}

	if height >= 5 {
		gui.Box(1, 1, width-2, 3, s)
		gui.Label(input, 2, 2, width-4, true, false, nil, s)
	}

	if height >= 6 {
		start, stop, sel := get()
		gui.List(files[start:stop], 2, 4, width-4, height-5, masks[start:stop], sel, s)
	}

	s.Sync() // due to bug in Show, blank cells don't update syle in vsCodeTerminal
	//s.Show()
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
	maybePanic := recover()
	s.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
	os.Exit(0)
}
