package main

import (
	"os"

	"golang.org/x/term"
)

func initScreen() *term.State {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	print("\033[?25l") //makes cursor invisible
	show("__",[]string{"blank","blank"})
	return oldState
}

func show(firstLine string,list []string) {
	print("\033[2K",firstLine,"\n") //clear line

	for i := 0; i < n; i++ {
		print("\033[2K") //clear line
		if i < len(list) {
			print(list[i])
		}
		print("\n")
	}

	for i := 0; i < n+1; i++ {
		print("\033[A") //move up
	}
}

func readInput() (byte, error) {
	buffer := make([]byte, 1)
	_, err := os.Stdin.Read(buffer)
	return buffer[0], err
}

func restore(oldState *term.State) {
	for i := 0; i < n+1; i++ {
		print("\033[2K\n") //clear line
	}
	for i := 0; i < n+1; i++ {
		print("\033[A") //move up
	}
	term.Restore(int(os.Stdin.Fd()), oldState)
}