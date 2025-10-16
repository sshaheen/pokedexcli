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
	const baseURL = "https://pokeapi.co/api/v2/location-area"
	scanner := bufio.NewScanner(os.Stdin)
	config := &models.Config{Next: baseURL, Previous: ""}

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
		command_str := cleaned_input[0]
		if command_str == "mapb" {
			if config.Previous == "" {
				fmt.Println("you're on the first page")
				continue
			} else {
				config.Next, config.Previous = config.Previous, config.Next
				command_str = "map"
			}
		}
		command, ok := command_map[command_str]
		if ok {
			command.Callback(config)
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
