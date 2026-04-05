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

type HistoryEntry struct {
	Command string
	Input   string
	Output  string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("STRING TRANSFORMER")
	fmt.Println("Command list: upper, lower, title, cap, snake, reverse, count, palindrome")
	fmt.Println("Type 'exit' to shut down program or 'history' for last 5 transformations")

	var history []HistoryEntry
	validCommands := map[string]bool{
		"upper": true, "lower": true, "cap": true,
		"title": true, "snake": true, "reverse": true,
		"history": true, "count": true, "palindrome": true,
	}

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}
		if strings.ToLower(input) == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye!")
			break
		}

		parts := strings.SplitN(input, " ", 2)
		command := strings.ToLower(parts[0])

		if command == "history" {
			if len(history) == 0 {
				fmt.Println("No history yet.")
				continue
			}
			fmt.Println("Last transformations:")
			for i, h := range history {
				fmt.Printf("%d) [%s]\n   Input: %s\n   Output: %s\n",
					i+1, h.Command, h.Input, h.Output)
			}
			continue
		}

		if len(parts) != 2 {
			fmt.Println("✗ No text provided. Example Usage: upper <text>")
			continue
		}

		text := parts[1]

		if !validCommands[command] {
			fmt.Printf("✗ Unknown command: \"%s\"\n", command)
			fmt.Println("Valid commands: upper, lower, cap, title, snake, reverse, palindrome, exit, count")
			continue
		}

		var result string

		switch command {
		case "palindrome":
			if IsPalindrome(text) {
				result = "is a palindrome!"
				fmt.Printf("✦ \"%s\" %s\n", text, result)
			} else {
				result = "is not a palindrome."
				fmt.Printf("✗ \"%s\" %s\n", text, result)
			}
		case "count":
			CountText(text)
			continue
		default:
			result = TransformString(text, command)
			fmt.Println("result:", result)
		}

		// Save to history
		history = append(history, HistoryEntry{
			Command: command,
			Input:   text,
			Output:  result,
		})

		// Keep only last 5 entries
		if len(history) > 5 {
			history = history[len(history)-5:]
		}
	}
}

// --- Transformations ---
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

func CountText(text string) {
	totalChars := len([]rune(text))
	letters := 0
	spaces := 0

	for _, r := range text {
		if unicode.IsLetter(r) {
			letters++
		}
		if unicode.IsSpace(r) {
			spaces++
		}
	}

	words := len(strings.Fields(text))

	fmt.Println("Total characters:", totalChars)
	fmt.Println("Total letters:", letters)
	fmt.Println("Total words:", words)
	fmt.Println("Total spaces:", spaces)
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
		return reverseText(text)
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

// --- Palindrome ---
func IsPalindrome(text string) bool {
	var cleaned []rune
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}
	return true
}