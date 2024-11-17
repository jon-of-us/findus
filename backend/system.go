package backend

import (
	"os"
	"strings"
)

type Path []string

func (p Path) String() string {
	return strings.Join(p, "/")
}

type File struct {
	Path     Path
	isFolder bool
}

func (f File) String() string {
	return f.Path.String()
}

type State struct {
	path        Path
	fuzzyFinder FuzzyFinder[File]
	searchDepth int
	lastFolders []Path
}

func InitState() *State {
	path := currentPath()
	return &State{path, FuzzyFinder[File]{}, 0, []Path{path}}
}
func (s *State) SetPath(path Path) {
	err := os.Chdir(path.String())
	if err == nil {
		s.path = path
		s.fuzzyFinder.Clear()
	}
}
func (s *State) PopPath() {
	if len(s.path) > 1 {
		s.SetPath(s.path[:len(s.path)-1])
	}
}
func (s *State) AddToPath(path Path) {
	s.SetPath(append(s.path, path...))
}

// func (s *State) FindMatches(query string, number int) ([]File, [][]bool) {
// 	matches, masks := s.fuzzyFinder.FindMatches(query, number)
// 	for len(matches) < number && s.searchDepth < config.MaxSearchDepth {
// 		s.searchDepth++
// 		for _, folder := range s.lastFolders {
// 			files, err := os.ReadDir(folder.String())
// 			if err != nil {
// 				continue
// 			}
// 			for _, file := range files {
// 				s.fuzzyFinder.AddObjects([]File{{folder, file.IsDir()}})
// 			}
// 		}
// 	}
// }
