package commands

import (
	"fmt"
	"os"

	"github.com/sshaheen/pokedexcli/internal/app"
)

func CommandExit(c *app.AppState) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
