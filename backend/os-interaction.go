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
		panic(err)
	}
	return strings.Split(path, PATH_SEPARATOR)
}

func (s *State) QuitAndSetPath() {
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
