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
	show([]rune{})
	return oldState
}

func show(input []rune) {
	print("\033[2K") //clear line
	print(string(input))
	print("\n")

	for i := 0; i < n; i++ {
		print("\033[2K") //clear line
		print("line ")
		print(i)
		print(" ")
		print(string(input))
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