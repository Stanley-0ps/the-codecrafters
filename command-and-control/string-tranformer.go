package main

import (
	"fmt"
	"strings"
)

// STRING TRANSFORMER

func handleStr(args []string) {
	if len(args) < 1 {
		fmt.Println("✗ Error: missing command.")
		return
	}

	cmd := strings.ToLower(args[0])

	if len(args) < 2 {
		fmt.Println("✗ No text provided.")
		return
	}

	text := strings.Join(args[1:], " ")

	var result string

	switch cmd {
	case "upper":
		result = strings.ToUpper(text)
	case "lower":
		result = strings.ToLower(text)
	case "title":
		result = Title(text)
	case "snake":
		result = Snake(text)
	case "reverse":
		result = reverseText(text)
	case "cap":
		result = Capitalize(text)
	default:
		fmt.Println("✗ Unknown string command.")
		return
	}

	fmt.Println("✦", result)
}
