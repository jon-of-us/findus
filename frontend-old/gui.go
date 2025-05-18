package main

var n int = 5

func main() {
	gui()
}

func gui() {
	oldState := initScreen()
	defer restore(oldState)

	input := make([]rune, 0)

mainloop:
	for {
		inputKey, err := readInput()
		if err != nil {
			panic(err)
		}

		switch inputKey {
		case 27: //win esc
			break mainloop
		case 127: //delete
			input = input[:0]
		default:
			input = append(input, rune(inputKey))
		}
		show(string(input), getListMock(string(input)))
	}
}

func getListMock(testWord string) []string {
	return []string{"jonas", "flo", "anton", testWord}
}
