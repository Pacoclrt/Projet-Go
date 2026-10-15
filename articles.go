package main

import "strings"

func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a" || words[i] == "A") && startsWithVowel(words[i+1]) {
			words[i] = words[i] + "n"
		}
	}
	return words
}

func startsWithVowel(word string) bool {
	first := strings.ToLower(word[:1])
	return first == "a" || first == "e" || first == "i" || first == "o" || first == "u" || first == "h"
}
