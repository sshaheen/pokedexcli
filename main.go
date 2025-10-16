package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sshaheen/pokedexcli/internal/commands"
	"github.com/sshaheen/pokedexcli/internal/models"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := &models.Config{}

	command_map := map[string]models.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commands.CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Provide user with usage info",
			Callback:    commands.CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Provide list of next 20 locations in Pokemon world",
			Callback:    commands.CommandMap,
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
			c_str.Callback(c)
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
