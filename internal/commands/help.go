package commands

import (
	"fmt"

	"github.com/sshaheen/pokedexcli/internal/app"
)

func CommandHelp(c *app.AppState) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Printf("\n\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays list of next 20 locations in Pokemon world")
	return nil
}
