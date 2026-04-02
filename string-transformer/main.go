// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: Stanley
// Squad:  ---

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("STRING TRANSFORMER")
	fmt.Println("Command list: upper, lower, title, cap, snake, reverse")
	fmt.Println("Type 'exit' to shut down program")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text()) // Trim whitespace

		if input == "" {
			// User pressed Enter with no text: show prompt again
			continue
		}

		if strings.ToLower(input) == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye!")
			break
		}

		parts := strings.SplitN(input, " ", 2)
		if len(parts) != 2 {
			fmt.Println("✗ No text provided. Example Usage: upper <text>")
			continue
		}

		command := strings.ToLower(parts[0])
		text := parts[1]

		validCommands := map[string]bool{
			"upper": true, "lower": true, "cap": true,
			"title": true, "snake": true, "reverse": true,
		}

		if !validCommands[command] {
			fmt.Printf("✗ Unknown command: \"%s\"\n", command)
			fmt.Println("Valid commands: upper, lower, cap, title, snake, reverse, exit")
			continue
		}

		result := TransformString(text, command)
		fmt.Println("result:", result)
	}
}

// Transformations
func Title(text string) string {
	if len(text) == 0 {
		return text
	}

	smallWords := map[string]bool{
		"a": true, "an": true, "the": true,
		"and": true, "but": true, "or": true,
		"for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true,
		"as": true, "is": true, "it": true,
	}

	words := strings.Fields(strings.ToLower(text))
	for i, word := range words {
		if i == 0 || !smallWords[word] {
			if len(word) > 0 {
				words[i] = strings.ToUpper(string(word[0])) + word[1:]
			}
		}
	}

	return strings.Join(words, " ")
}

func reverse(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		runes := []rune(word)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
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

func TransformString(text, mode string) string {
	switch mode {
	case "upper":
		return strings.ToUpper(text)
	case "lower":
		return strings.ToLower(text)
	case "title":
		return Title(text)
	case "reverse":
		return reverse(text)
	case "snake":
		return Snake(text)
	case "cap":
		words := strings.Fields(text)
		for i, word := range words {
			if len(word) > 0 {
				words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			}
		}
		return strings.Join(words, " ")
	default:
		return text
	}
}
