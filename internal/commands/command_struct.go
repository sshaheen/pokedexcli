package commands

import "github.com/sshaheen/pokedexcli/internal/app"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(a *app.AppState) error
}
