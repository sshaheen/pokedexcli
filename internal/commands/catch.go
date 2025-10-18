package commands

import (
	"fmt"
	"math/rand/v2"

	"github.com/sshaheen/pokedexcli/internal/app"
	"github.com/sshaheen/pokedexcli/internal/models"
	"github.com/sshaheen/pokedexcli/internal/pokeapi"
)

func CommandCatch(state *app.AppState) error {
	pokemon, err := pokeapi.Fetch[models.Pokemon](state.Client, state.Config.Next)
	if err != nil {
		fmt.Printf("%s was not found...\n", state.TargetPokemon)
		return err
	}

	if _, ok := state.Pokedex[pokemon.Name]; ok {
		fmt.Printf("%s is already in your Pokedex.\n", pokemon.Name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	bExp := float64(pokemon.BaseExp)

	if bExp < 50 {
		bExp = 50
	} else if bExp > 300 {
		bExp = 300
	}

	chance := 0.75 - (float64(bExp)-50)/250*0.5
	caught := rand.Float64() < chance

	if caught {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		state.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
