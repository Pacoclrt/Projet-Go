package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Utilisation : go run . sample.txt result.txt")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur : impossible de lire", os.Args[1])
		return
	}
	result := correctText(string(data))
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err != nil {
		fmt.Println("Erreur : impossible d'écrire", os.Args[2])
	}
}

func correctText(text string) string {
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = correctLine(lines[i])
	}
	return strings.Join(lines, "\n")
}

func correctLine(line string) string {
	words := strings.Fields(line)
	words = applyTags(words)
	words = fixPunctuation(words)
	words = fixQuotes(words)
	words = fixArticles(words)
	return strings.Join(words, " ")
}
