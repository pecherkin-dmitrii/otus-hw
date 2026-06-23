package hw03frequencyanalysis

import (
	"cmp"
	"slices"
	"strings"
)

const amountOfTopWordsToTake = 10

type wordWithAmount struct {
	word   string
	amount int
}

var wordWithAmountComparator = func(a, b wordWithAmount) int {
	if n := cmp.Compare(b.amount, a.amount); n != 0 {
		return n
	}
	return cmp.Compare(a.word, b.word)
}

func Top10(str string) []string {
	words := strings.Fields(str)

	uniqueWordsWithAmount := make(map[string]int)
	for _, word := range words {
		uniqueWordsWithAmount[word]++
	}

	wordsWithAmount := make([]wordWithAmount, 0, len(uniqueWordsWithAmount))
	for word, amount := range uniqueWordsWithAmount {
		wordWithAmount := wordWithAmount{
			word:   word,
			amount: amount,
		}
		wordsWithAmount = append(wordsWithAmount, wordWithAmount)
	}
	slices.SortStableFunc(wordsWithAmount, wordWithAmountComparator)

	result := make([]string, 0, amountOfTopWordsToTake)
	for i, w := range wordsWithAmount {
		if i == cap(result) {
			break
		}
		result = append(result, w.word)
	}

	return result
}
