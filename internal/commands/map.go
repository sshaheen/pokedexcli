package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sshaheen/pokedexcli/internal/models"
)

func CommandMap(c *models.Config) error {
	res, err := http.Get(c.Next)

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

	var config models.Config

	err = json.Unmarshal(body, &config)

	if err != nil {
		return err
	}

	c.Next = config.Next
	c.Previous = config.Previous

	for _, item := range map_data.Results {
		fmt.Printf("%s\n", item.Name)
	}

	return nil
}
