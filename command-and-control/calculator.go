package main

import (
	"fmt"
	"strings"
)

// CALCULATOR

func handleCalc(args []string) {
	if len(args) < 1 {
		fmt.Println("✗ Error: missing operation.")
		return
	}

	op := strings.ToLower(args[0])

	if op == "last" {
		fmt.Println("✦ Result:", lastResult)
		return
	}

	if len(args) < 3 {
		fmt.Println("✗ Error: missing arguments.")
		return
	}

	a, err1 := parseNumber(args[1])
	b, err2 := parseNumber(args[2])

	if err1 != nil || err2 != nil {
		fmt.Println("✗ Error: invalid number.")
		return
	}

	var result float64

	switch op {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			fmt.Println("✗ Error: cannot divide by zero.")
			return
		}
		result = a / b
	case "mod":
		if b == 0 {
			fmt.Println("✗ Error: cannot mod by zero.")
			return
		}
		result = float64(int(a) % int(b))
	case "pow":
		result = pow(a, b)
	default:
		fmt.Println("✗ Unknown calc command.")
		return
	}

	lastResult = result
	fmt.Println("✦ Result:", result)
}

func pow(a, b float64) float64 {
	result := 1.0
	for i := 0; i < int(b); i++ {
		result *= a
	}
	return result
}
