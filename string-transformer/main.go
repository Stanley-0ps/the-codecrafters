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
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("String transformer")
	fmt.Println("Command list: Upper, lower, title, cap, snake, reverse")
	fmt.Println("Type 'quit' to shut down program")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if Strings.ToLower(input) == "quit" {
			fmt.Println("Shutting down transformer")
			break
		}
		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Unknown command. Example Usage: upper <text>")
			continue
		}
		command := strings.ToLower(parts[0])
		text := parts[1]

		if command != "upper" || command != "lower" || command != "cap" || command != "title" || command != "snake" || command != "reverse" {
			fmt.Println("Unknown command. Try again!")
			continue
		}

	}
}

func TransformString(text string) {
	switch {
	case "upper":
		{
			return strings.ToUpper(text)
		}
	case "lower":
		{
			return strings.ToLower(text)
		}
	case "cap":
		{
			return strings.Title(text)
		}
	}
}
