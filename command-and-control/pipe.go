package main

import (
	"fmt"
	"strings"
)

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
