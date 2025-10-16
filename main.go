package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	Next     string
	Previous string
}

type mapData struct {
	Results []location
}

type location struct {
	Name string
	Url  string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := &config{}

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
		"map": {
			name:        "map",
			description: "Provide list of next 20 locations in Pokemon world",
			callback:    commandMap,
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
			c_str.callback(c)
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

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Printf("\n\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays list of next 20 locations in Pokemon world")
	return nil
}

func commandMap(c *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	res.Body.Close()

	var map_data mapData

	err = json.Unmarshal(body, &map_data)

	if err != nil {
		return err
	}

	for _, item := range map_data.Results {
		fmt.Printf("Name: %s\n", item.Name)
	}

	return nil
}
