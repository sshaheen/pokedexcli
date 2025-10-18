package commands

import (
	"fmt"

	"github.com/sshaheen/pokedexcli/internal/app"
)

func CommandInspect(state *app.AppState) error {
	target_pokemon := state.TargetPokemon
	pokemon, ok := state.Pokedex[target_pokemon]
	if !ok {
		fmt.Printf("%s does not exist in your Pokedex...\n", target_pokemon)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, p_type := range pokemon.Types {
		fmt.Printf("  -%s: %d\n", p_type.Type.Name, p_type.Slot)
	}

	return nil
}
