package main

import "testing"

func check(t *testing.T, input string, expected string) {
	result := correctText(input)
	if result != expected {
		t.Errorf("\nentrée  : %q\nobtenu  : %q\nattendu : %q", input, result, expected)
	}
}

func TestCase(t *testing.T) {
	check(t, "Ready, set, go (up) !", "Ready, set, GO!")
	check(t, "I should stop SHOUTING (low)", "I should stop shouting")
	check(t, "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge")
	check(t, "This is so exciting (up, 2)", "This is SO EXCITING")
	check(t, "this is so exciting (cap, 3)", "this Is So Exciting")
}

func TestNumbers(t *testing.T) {
	check(t, "1E (hex) files were added", "30 files were added")
	check(t, "It has been 10 (bin) years", "It has been 2 years")
}

func TestPunctuation(t *testing.T) {
	check(t, "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!")
	check(t, "I was thinking ... You were right", "I was thinking... You were right")
}

func TestQuotes(t *testing.T) {
	check(t, "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'")
	check(t, "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'")
}

func TestArticles(t *testing.T) {
	check(t, "There it was. A amazing rock!", "There it was. An amazing rock!")
	check(t, "There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you.")
}

func TestExamples(t *testing.T) {
	check(t, "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom", "It was the best of times, it was the worst of TIMES, it was the age of wisdom")
	check(t, "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68.")
	check(t, "Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?")
}

func TestStrangeCases(t *testing.T) {
	check(t, "", "")
	check(t, "(up) hello", "hello")
	check(t, "hello world (cap, 6)", "Hello World")
	check(t, "zz (hex) is not a number", "zz is not a number")
}
