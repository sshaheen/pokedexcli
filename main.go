package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	command_map := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Provide user with usage info",
			callback:    commandHelp,
		},
	}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input_text := scanner.Text()
		cleaned_input := cleanInput(input_text)
		command := cleaned_input[0]
		c_str, ok := command_map[command]
		if ok {
			c_str.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	trimmed_text := strings.TrimSpace(text)
	words := strings.Split(trimmed_text, " ")
	var result []string
	for i := range words {
		if words[i] != "" {
			result = append(result, strings.TrimSpace(strings.ToLower(words[i])))
		}
	}
	return result
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Printf("\n\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}
