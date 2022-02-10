package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Freq struct {
	word  string
	count int
}

func HasFreq(word string, freqs []Freq) (bool, int) {
	for i, f := range freqs {
		if f.word == word {
			return true, i
		}
	}

	return false, 0
}

func GetWords(freqs []Freq) []string {
	var words []string

	for _, f := range freqs {
		words = append(words, f.word)
	}

	return words
}

func Top10(text string) []string {
	freqs := make([]Freq, 0)
	words := strings.Fields(text)

	if len(words) == 0 {
		return nil
	}

	for _, w := range words {
		h, i := HasFreq(w, freqs)
		if h {
			freqs[i].count++
		} else {
			freqs = append(freqs, Freq{w, 1})
		}
	}

	sort.Slice(freqs, func(i, j int) bool {
		if freqs[i].count == freqs[j].count {
			return freqs[i].word < freqs[j].word
		}

		return freqs[i].count > freqs[j].count
	})

	return GetWords(freqs[:10])
}
