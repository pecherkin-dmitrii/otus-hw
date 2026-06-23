package hw03frequencyanalysis

import (
	"cmp"
	"slices"
	"strings"
)

const amountOfTopWordsToTake = 10

type wordWithCount struct {
	word  string
	count int
}

var wordWithCountComparator = func(a, b wordWithCount) int {
	if n := cmp.Compare(b.count, a.count); n != 0 {
		return n
	}
	return cmp.Compare(a.word, b.word)
}

func Top10(str string) []string {
	splitted := strings.Fields(str)

	wordsMap := make(map[string]int)
	for _, word := range splitted {
		wordsMap[word]++
	}

	wordsSlice := make([]wordWithCount, 0, len(wordsMap))
	for word, count := range wordsMap {
		wordWithCount := wordWithCount{
			word:  word,
			count: count,
		}
		wordsSlice = append(wordsSlice, wordWithCount)
	}
	slices.SortStableFunc(wordsSlice, wordWithCountComparator)

	result := make([]string, 0, amountOfTopWordsToTake)
	for i, w := range wordsSlice {
		if i == cap(result) {
			break
		}
		result = append(result, w.word)
	}

	return result
}
