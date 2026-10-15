package main

func fixQuotes(words []string) []string {
	result := []string{}
	open := false
	for i := 0; i < len(words); i++ {
		if words[i] != "'" {
			result = append(result, words[i])
		} else if !open && i+1 < len(words) {
			words[i+1] = "'" + words[i+1]
			open = true
		} else if open && len(result) > 0 {
			result[len(result)-1] += "'"
			open = false
		} else {
			result = append(result, words[i])
		}
	}
	return result
}
