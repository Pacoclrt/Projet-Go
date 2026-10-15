package main

func isPunctuation(c byte) bool {
	return c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';'
}

func fixPunctuation(words []string) []string {
	result := []string{}
	for _, word := range words {
		for len(word) > 0 && len(result) > 0 && isPunctuation(word[0]) {
			result[len(result)-1] += string(word[0])
			word = word[1:]
		}
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}
