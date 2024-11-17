package backend

type FuzzyFinder struct {
	words []string
}

func (f *FuzzyFinder) AddWords(newWords []string) {
	f.words = append(f.words, newWords...)
}
func (f *FuzzyFinder) ClearWords() {
	f.words = []string{}
}
func (f *FuzzyFinder) FindMatches(query string) ([]string, [][]bool) {
	var matches []string
	var masks [][]bool
	for _, word := range f.words {
		queryIdx := 0
		wordIdx := 0
		mask := make([]bool, len(word))
		for queryIdx < len(query) && wordIdx < len(word) {
			if query[queryIdx] == word[wordIdx] {
				queryIdx++
				mask[wordIdx] = true
			}
			wordIdx++
		}
		if queryIdx == len(query) {
			matches = append(matches, word)
			masks = append(masks, mask)
		}
	}
	return matches, masks
}
