package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := "Lorem, lorem ipsum"
	num, occurrence := countWords(numbers)
	fmt.Println(num, occurrence)
}

func splitSentence(r rune) bool {
    return r == ',' || r == '.' || r == ' '
}

func countWords(sentence string) (int, map[string]int) {
	sentenceSlice := strings.FieldsFunc(sentence, splitSentence)
	wordCounts := make(map[string]int)

	for _, word := range sentenceSlice {

		word = strings.ToLower(word)
		_, exists := wordCounts[word]
		
		if exists {
			wordCounts[word] += 1
		} else {
			wordCounts[word] = 1
		}
	}

	return len(wordCounts), wordCounts
}