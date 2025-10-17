package commands

import (
	"fmt"

	"github.com/sshaheen/pokedexcli/internal/app"
)

func CommandMap(state *app.AppState) error {
	data, config, err := state.Client.Fetch(state.Config.Next)

	if err != nil {
		return err
	}

	state.Config.Next = config.Next
	state.Config.Previous = config.Previous

	for _, item := range data.Results {
		fmt.Printf("%s\n", item.Name)
	}

	return nil
}
