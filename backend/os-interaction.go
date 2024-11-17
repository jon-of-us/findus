//go:build windows

package backend

import (
	"fmt"
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
	fmt.Println(path)
	return strings.Split(path, PATH_SEPARATOR)
}

func (s *System) QuitAndSetPath() {
	pathString := strings.Join(s.Path, PATH_SEPARATOR)
	cmd := exec.Command("cmd", "/C", "cd", "/d", pathString,
		"&&", "powershell", "-NoLogo")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return
	}
}
