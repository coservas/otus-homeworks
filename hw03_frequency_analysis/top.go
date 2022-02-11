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
	words := make([]string, len(freqs))

	for i, f := range freqs {
		words[i] = f.word
	}

	return words
}

func Top10(text string) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return nil
	}

	var freqs []Freq
	for _, w := range words {
		has, i := HasFreq(w, freqs)
		if has {
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

	var length int
	if len(freqs) > 10 {
		length = 10
	} else {
		length = len(freqs)
	}

	return GetWords(freqs[:length])
}
