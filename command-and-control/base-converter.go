package main

import (
	"fmt"
	"strconv"
	"strings"
)

// BASE CONVERTER

func handleBase(args []string) {
	if len(args) < 1 {
		fmt.Println("✗ Error: missing base command.")
		return
	}

	cmd := strings.ToLower(args[0])

	if len(args) < 2 {
		fmt.Println("✗ Error: missing number.")
		return
	}

	input := args[1]

	if input == "last" {
		input = fmt.Sprintf("%.0f", lastResult)
	}

	switch cmd {

	case "hex":
		val, err := strconv.ParseInt(input, 16, 64)
		if err != nil {
			fmt.Printf("✗ Error: %s is not valid hex.\n", input)
			return
		}
		lastResult = float64(val)
		fmt.Println("✦ Decimal:", val)

	case "bin":
		val, err := strconv.ParseInt(input, 2, 64)
		if err != nil {
			fmt.Printf("✗ Error: %s is not valid binary.\n", input)
			return
		}
		lastResult = float64(val)
		fmt.Println("✦ Decimal:", val)

	case "dec":
		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("✗ Error: invalid decimal number.")
			return
		}
		lastResult = float64(val)
		fmt.Println("✦ Binary :", strconv.FormatInt(int64(val), 2))
		fmt.Println("✦ Hex    :", strings.ToUpper(strconv.FormatInt(int64(val), 16)))

	default:
		fmt.Println("✗ Error: unknown base command.")
	}
}
