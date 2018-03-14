package corpus

import (
	"regexp"
	"sort"
	"strings"
)

// WordFrequency keeps track of Word and number of occurances
type WordFrequency struct {
	Word  string
	Count int
}

// Analyze does something, Chris told me to write it this way!
func Analyze(rawSt string) []WordFrequency {
	sanitizedStArr := sanitize(rawSt)
	wordFreqs := wordCountGenerator(sanitizedStArr)
	return descSortOnWordCount(wordFreqs)
}

func sanitize(rawSt string) []string {
	rg := regexp.MustCompile("[^a-zA-Z0-9' ]+")
	str := rg.ReplaceAllLiteralString(rawSt, "")
	return strings.Fields(strings.ToLower(str))
}

func wordCountGenerator(sanitizedStArr []string) []WordFrequency {
	wordMap := make(map[string]int)

	for _, word := range sanitizedStArr {
		wordMap[word]++
	}
	wordMapLen := len(wordMap)
	wordFreqs := make([]WordFrequency, wordMapLen)

	for word, count := range wordMap {
		wordMapLen--
		wordFreqs[wordMapLen] = WordFrequency{Word: word, Count: count}
	}

	return wordFreqs
}

func descSortOnWordCount(wordFreqs []WordFrequency) []WordFrequency {
	sort.Slice(wordFreqs, func(i, j int) bool { return wordFreqs[i].Count > wordFreqs[j].Count })
	return wordFreqs
}
