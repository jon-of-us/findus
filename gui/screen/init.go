package screen

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

func Init() tcell.Screen {
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

func Quit(s tcell.Screen) {
	s.Fini()
	os.Exit(0)
}
