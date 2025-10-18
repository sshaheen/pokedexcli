package commands

import (
	"fmt"

	"github.com/sshaheen/pokedexcli/internal/app"
)

func CommandPokedex(state *app.AppState) error {
	if len(state.Pokedex) == 0 {
		fmt.Println("Try catching some Pokemon first!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range state.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
