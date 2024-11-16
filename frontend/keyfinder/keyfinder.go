package main

import (
	"os"

	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1)
	_, err2 := os.Stdin.Read(buffer)
	if err2 != nil {
		panic(err2)
	}
	print(buffer[0])
}