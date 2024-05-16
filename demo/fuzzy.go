package main

func fzFind(pattern []rune, files [][]rune) (matches [][]rune, masks [][]bool) {
	for _, file := range files {
		if match, mask := fzCheck(pattern, file); match {
			matches = append(matches, file)
			masks = append(masks, mask)
		}
	}
	return
}

func fzCheck(pattern, word []rune) (bool, []bool) {
	pId, mask := 0, make([]bool, len(word))
	for wId := 0; pId < len(pattern) && wId < len(word); wId++ {
		if pattern[pId] == word[wId] {
			pId++
			mask[wId] = true
		}
	}
	return pId == len(pattern), mask
}
