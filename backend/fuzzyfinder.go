package backend

import "fmt"

type FuzzyFinder[T fmt.Stringer] struct {
	objects []T
	strings []string
}

func (f *FuzzyFinder[T]) AddObjects(newObjects []T) {
	strings := make([]string, len(newObjects))
	for i, obj := range newObjects {
		strings[i] = obj.String()
	}
	f.objects = append(f.objects, newObjects...)
	f.strings = append(f.strings, strings...)
}
func (f *FuzzyFinder[T]) Clear() {
	f.strings = []string{}
}
func (f *FuzzyFinder[T]) FindMatches(query string, number int) ([]T, [][]bool) {
	var matches []T
	var masks [][]bool
	n := 0
	for i, str := range f.strings {
		if n >= number {
			break
		}
		queryIdx := 0
		strIdx := 0
		mask := make([]bool, len(str))
		for queryIdx < len(query) && strIdx < len(str) && n < number {
			if query[queryIdx] == str[strIdx] {
				queryIdx++
				mask[strIdx] = true
			}
			strIdx++
		}
		if queryIdx == len(query) {
			matches = append(matches, f.objects[i])
			masks = append(masks, mask)
			n++
		}
	}
	return matches, masks
}
