package main

import (
	"strconv"
	"strings"
	"unicode"
)

func applyTags(words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "(hex)" {
			convertNumber(result, 16)
		} else if word == "(bin)" {
			convertNumber(result, 2)
		} else if word == "(up)" || word == "(low)" || word == "(cap)" {
			mode := word[1 : len(word)-1]
			changeCase(result, mode, 1)
		} else if isNumberTag(word) && i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
			mode := word[1 : len(word)-1]
			number := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(number)
			if err == nil {
				changeCase(result, mode, count)
				i++
			} else {
				result = append(result, word)
			}
		} else {
			result = append(result, word)
		}
	}
	return result
}

func isNumberTag(word string) bool {
	return word == "(up," || word == "(low," || word == "(cap,"
}

func convertNumber(words []string, base int) {
	if len(words) == 0 {
		return
	}
	last := len(words) - 1
	number, err := strconv.ParseInt(words[last], base, 64)
	if err == nil {
		words[last] = strconv.FormatInt(number, 10)
	}
}

func changeCase(words []string, mode string, count int) {
	start := len(words) - count
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		if mode == "up" {
			words[i] = strings.ToUpper(words[i])
		} else if mode == "low" {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = capitalize(words[i])
		}
	}
}

func capitalize(word string) string {
	letters := []rune(strings.ToLower(word))
	if len(letters) > 0 {
		letters[0] = unicode.ToUpper(letters[0])
	}
	return string(letters)
}
