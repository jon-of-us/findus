package backend

import (
	"findus/config"
	"os"
	"strings"
)

type Path []string

func (p Path) String() string {
	return strings.Join(p, "/")
}

type File struct {
	Path     Path
	IsFolder bool
}

func (f File) String() string {
	return f.Path.String()
}

type State struct {
	Path        Path
	fuzzyFinder FuzzyFinder[File]
	searchDepth int
	// openSubfoldersRelative contains all subfolders used by the fuzzy find as a relative path
	openSubfoldersRelative []Path
	// ForwardPath last Element always is the previous folder, if its deeper than the current folder
	ForwardPath Path
}

func NewState() *State {
	return &State{
		Path:                   currentPath(),
		openSubfoldersRelative: []Path{{}},
	}
}
func (s *State) setPath(path Path) {
	err := os.Chdir(path.String())
	if err == nil {
		s.Path = path
		s.fuzzyFinder.Clear()
		s.openSubfoldersRelative = []Path{{}}
		s.searchDepth = 0
	}
}
func (s *State) PopPath() {
	if len(s.Path) > 1 {
		s.ForwardPath = append(s.ForwardPath, s.Path[len(s.Path)-1])
		s.setPath(s.Path[:len(s.Path)-1])
	}
}
func (s *State) AddToPath(path Path) {
	s.setPath(append(s.Path, path...))
	for len(path) > 0 && len(s.ForwardPath) > 0 {
		if path[0] == s.ForwardPath[len(s.ForwardPath)-1] {
			s.ForwardPath = s.ForwardPath[:len(s.ForwardPath)-1]
			path = path[1:]
		} else {
			s.ForwardPath = []string{}
		}
	}
}
func (s *State) GoToPath(path Path) {
	s.setPath(path)
	s.ForwardPath = []string{}
}

func (s *State) FindMatches(query string, number int) ([]File, [][]bool) {
	matches, masks := s.fuzzyFinder.FindMatches(query, number)
	for len(matches) < number && s.searchDepth < config.MaxSearchDepth {
		s.searchDepth++
		subfolders := []Path{}
		subfiles := []File{}
		for _, folder := range s.openSubfoldersRelative {
			absoluteFolder := append(s.Path, folder...)
			files, err := os.ReadDir(absoluteFolder.String())
			if err != nil {
				continue
			}
			for _, file := range files {
				fileObj := File{append(folder, file.Name()), file.IsDir()}
				subfiles = append(subfiles, fileObj)
				if file.IsDir() {
					subfolders = append(subfolders, fileObj.Path)
				}
			}
		}
		s.fuzzyFinder.AddObjects(subfiles)
		s.openSubfoldersRelative = subfolders
		matches, masks = s.fuzzyFinder.FindMatches(query, number)
	}
	return matches, masks
}
