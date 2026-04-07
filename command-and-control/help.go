package main

import "fmt"

func showHistory() {
	for i, h := range history {
		fmt.Printf("%d. %s\n", i+1, h)
	}
}

//HELP / HISTORY

func clearSession() {
	fmt.Print("Type CONFIRM to clear the session: ")
	var input string
	fmt.Scanln(&input)

	if input == "CONFIRM" {
		history = []string{}
		lastResult = 0
		fmt.Println("✔ Session cleared.")
	} else {
		fmt.Println("Cancelled.")
	}
}

func showHelp() {
	fmt.Println("Available Commands:")
	fmt.Println(" calc add|sub|mul|div|mod|pow <a> <b>")
	fmt.Println(" base dec|hex|bin <number>")
	fmt.Println(" str upper|lower|title|snake|reverse|cap <text>")
	fmt.Println(" history, clear, exit")
}
