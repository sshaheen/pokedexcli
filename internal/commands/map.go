package commands

import (
	"fmt"

	"github.com/sshaheen/pokedexcli/internal/app"
	"github.com/sshaheen/pokedexcli/internal/models"
	"github.com/sshaheen/pokedexcli/internal/pokeapi"
)

func CommandMap(state *app.AppState) error {
	data, err := pokeapi.Fetch[models.MapData](state.Client, state.Config.Next)

	if err != nil {
		return err
	}

	state.Config.Next = data.Next
	state.Config.Previous = data.Previous

	for _, item := range data.Results {
		fmt.Printf("%s\n", item.Name)
	}

	return nil
}
