// CodeCrafters — Hackathon 002
// Squad: nil
// Members: nil

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lastResult float64
var history []string

func main() {
	fmt.Println("═══════════════════════════════════════════════")
	fmt.Println("  SENTINEL — COMMAND & CONTROL CONSOLE")
	fmt.Println("     All systems nominal. Type 'help' to begin.")
	fmt.Println("═══════════════════════════════════════════════")

	startConsole()
}

func startConsole() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("C&C> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		// Save history
		history = append(history, input)
		if len(history) > 10 {
			history = history[1:]
		}

		handleInput(input)
	}
}

func handleInput(input string) {
	if strings.Contains(input, "|") {
		handlePipe(input)
		return
	}

	parts := strings.Fields(input)
	cmd := strings.ToLower(parts[0])

	switch cmd {
	case "calc":
		handleCalc(parts[1:])
	case "base":
		handleBase(parts[1:])
	case "str":
		handleStr(parts[1:])
	case "history":
		showHistory()
	case "clear":
		clearSession()
	case "help":
		showHelp()
	case "exit":
		fmt.Println("Shutting down...")
		os.Exit(0)
	default:
		fmt.Println("✗ Unknown command.")
	}
}

// PIPE SYSTEM

func handlePipe(input string) {
	parts := strings.Split(input, "|")

	if len(parts) != 2 {
		fmt.Println("✗ Error: invalid pipe usage.")
		return
	}

	left := strings.TrimSpace(parts[0])
	right := strings.TrimSpace(parts[1])

	executeAndStore(left)

	newCommand := right + " " + fmt.Sprintf("%.0f", lastResult)
	handleInput(newCommand)
}

func executeAndStore(cmd string) {
	parts := strings.Fields(cmd)

	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "calc":
		handleCalc(parts[1:])
	case "base":
		handleBase(parts[1:])
	}
}
