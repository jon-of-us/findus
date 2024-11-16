package main

var n int = 5

func main() {
	gui()
}

func gui() {
	oldState := initScreen()
	defer restore(oldState)

	input := make([]rune, 0)

	for loop := true; loop; {
		inputKey, err := readInput()
		if err != nil {
			panic(err)
		}

		switch inputKey {
		case 27: //win esc
			loop = false
		default:
			input = append(input, rune(inputKey))
			show(input)
		}
	}
}
