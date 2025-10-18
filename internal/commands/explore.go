package commands

import (
	"fmt"

	"github.com/sshaheen/pokedexcli/internal/app"
	"github.com/sshaheen/pokedexcli/internal/models"
	"github.com/sshaheen/pokedexcli/internal/pokeapi"
)

func CommandExplore(state *app.AppState) error {
	data, err := pokeapi.Fetch[models.AreaDetail](state.Client, state.Config.Next)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", data.Name)

	for _, item := range data.PokemonEncounters {
		fmt.Println(item.Pokemon.Name)
	}

	return nil
}
