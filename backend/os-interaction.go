//go:build windows

package backend

import (
	"os"
	"os/exec"
	"strings"
)

const PATH_SEPARATOR = "\\"

func currentPath() []string {
	path, err := os.Getwd()
	if err != nil {
		panic("Failed to get current working directory: " + err.Error())
	}
	return strings.Split(path, PATH_SEPARATOR)
}

// Exit changes the current working directory to the path stored in the State
func (s *State) Exit() {
	cmd := exec.Command("cmd", "/C", "cd", "/d", s.Path.String(),
		"&&", "powershell", "-NoLogo")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return
	}
}
