package main

import (
	"strconv"
	"strings"
	"unicode"
)

// HELPERS

func parseNumber(input string) (float64, error) {
	if input == "last" {
		return lastResult, nil
	}
	return strconv.ParseFloat(input, 64)
}

func Title(text string) string {
	words := strings.Fields(strings.ToLower(text))
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

func reverseText(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Snake(text string) string {
	text = strings.ToLower(text)
	var b strings.Builder
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		} else if unicode.IsSpace(r) {
			b.WriteRune('_')
		}
	}
	return b.String()
}

func Capitalize(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}
