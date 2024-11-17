package main

import "findus/backend"

func main() {
	system := backend.InitSystem()
	system.AddToPath([]string{"not working"})

	system.AddToPath([]string{"test"})
	system.QuitAndSetPath()
}
