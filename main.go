package main

import "findus/backend"

func main() {
	bState := backend.InitState()
	bState.AddToPath(backend.Path{"test"})
	bState.PopPath()
	bState.AddToPath(backend.Path{".git"})
	matches, masks := bState.FindMatches("eam", 10)
	for i, match := range matches {
		println(match.String())
		for _, mask := range masks[i] {
			if mask {
				print("1")
			} else {
				print("0")
			}
		}
		println()
	}

}
