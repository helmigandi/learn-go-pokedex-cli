package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// This blocks the code and waits for input, once the user types something and presses enter
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

// The purpose of this function will be to split the users input into "words" based on whitespace.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}
