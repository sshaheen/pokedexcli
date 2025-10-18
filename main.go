package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sshaheen/pokedexcli/internal/app"
	"github.com/sshaheen/pokedexcli/internal/commands"
	"github.com/sshaheen/pokedexcli/internal/models"
	"github.com/sshaheen/pokedexcli/internal/pokeapi"
	"github.com/sshaheen/pokedexcli/internal/pokecache"
)

func main() {
	const baseURL = "https://pokeapi.co/api/v2/location-area"
	const pokemonURL = "https://pokeapi.co/api/v2/pokemon"
	scanner := bufio.NewScanner(os.Stdin)
	config := &models.Config{Next: baseURL, Previous: ""}
	cache := pokecache.NewCache(5 * time.Second)
	client := pokeapi.NewClient(cache)
	pokedex := make(map[string]models.Pokemon)
	state := &app.AppState{Config: config, Client: client, Pokedex: pokedex}

	command_map := map[string]commands.CliCommand{
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
		"explore": {
			Name:        "explore",
			Description: "List Pokemon in area",
			Callback:    commands.CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a Pokemon",
			Callback:    commands.CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Get details on Pokemon in Pokedex",
			Callback:    commands.CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Print names of all Pokemon in Pokedex",
			Callback:    commands.CommandPokedex,
		},
	}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input_text := scanner.Text()
		cleaned_input := cleanInput(input_text)
		if len(cleaned_input) == 2 {
			command_str := cleaned_input[0]
			switch command_str {
			case "explore":
				area := cleaned_input[1]
				command, ok := command_map[command_str]
				if ok {
					state.Config.Next = fmt.Sprintf("%s/%s", baseURL, area)
					command.Callback(state)
				} else {
					fmt.Println("Unknown command")
				}
			case "catch":
				pokemon := cleaned_input[1]
				command, ok := command_map[command_str]
				if ok {
					state.Config.Next = fmt.Sprintf("%s/%s", pokemonURL, pokemon)
					state.TargetPokemon = pokemon
					command.Callback(state)
				} else {
					fmt.Println("Unknown command")
				}
			case "inspect":
				pokemon := cleaned_input[1]
				command, ok := command_map[command_str]
				if ok {
					state.TargetPokemon = pokemon
					command.Callback(state)
				} else {
					fmt.Printf("Unknown command")
				}
			}
		} else {
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
				command.Callback(state)
			} else {
				fmt.Println("Unknown command")
			}
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
