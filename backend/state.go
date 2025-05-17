package backend

import (
	"findus/config"
	"os"
	"reflect"
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
	Path        Path
	fuzzyFinder FuzzyFinder[File]
	searchDepth int
	// openSubfoldersRelative contains empty path for current Folder
	openSubfoldersRelative []Path
	// forwardPath last Element always is the previous folder, if its deeper than the current folder
	forwardPath []string
}

func InitState() *State {
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
	}
}
func (s *State) PopPath() {
	if len(s.Path) > 1 {
		s.forwardPath = append(s.forwardPath, s.Path[len(s.Path)-1])
		s.setPath(s.Path[:len(s.Path)-1])

	}
}
func (s *State) AddToPath(path Path) {
	s.setPath(append(s.Path, path...))
	if len(s.forwardPath) != 0 {
		if reflect.DeepEqual(path, s.forwardPath[len(s.forwardPath)-1]) {
			s.forwardPath = s.forwardPath[:len(s.forwardPath)-1]
		}
	}
}
func (s *State) GoToPath(path Path) {
	s.setPath(path)
	s.forwardPath = []string{}
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
