package app

import (
	"github.com/sshaheen/pokedexcli/internal/models"
	"github.com/sshaheen/pokedexcli/internal/pokeapi"
)

type AppState struct {
	Config *models.Config
	Client *pokeapi.Client
}
