package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sshaheen/pokedexcli/internal/models"
)

func CommandMap(c *models.Config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	res.Body.Close()

	var map_data models.MapData

	err = json.Unmarshal(body, &map_data)

	if err != nil {
		return err
	}

	for _, item := range map_data.Results {
		fmt.Printf("%s\n", item.Name)
	}

	return nil
}
