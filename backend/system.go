package backend

import (
	"os"
	"strings"
)

type Path []string

func (p Path) String() string {
	return strings.Join(p, "/")
}

type System struct {
	Path Path
}

func InitSystem() *System {
	return &System{currentPath()}
}
func (s *System) SetPath(path Path) {
	pathString := strings.Join(path, PATH_SEPARATOR)
	err := os.Chdir(pathString)
	if err == nil {
		s.Path = path
	}
}
func (s *System) PopPath() {
	if len(s.Path) > 1 {
		s.SetPath(s.Path[:len(s.Path)-1])
	}
}
func (s *System) AddToPath(path Path) {
	s.SetPath(append(s.Path, path...))
}
