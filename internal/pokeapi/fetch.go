package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sshaheen/pokedexcli/internal/models"
)

func (c *Client) Fetch(url string) (models.MapData, models.Config, error) {
	if raw_data, ok := c.cache.Get(url); ok {
		var map_data models.MapData

		err := json.Unmarshal(raw_data, &map_data)
		if err != nil {
			return models.MapData{}, models.Config{}, err
		}

		var config models.Config

		err = json.Unmarshal(raw_data, &config)

		if err != nil {
			return models.MapData{}, models.Config{}, err
		}

		return map_data, config, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return models.MapData{}, models.Config{}, err
	}

	body, err := io.ReadAll(res.Body)

	c.cache.Add(url, body)

	if err != nil {
		return models.MapData{}, models.Config{}, err
	}

	res.Body.Close()

	var map_data models.MapData

	err = json.Unmarshal(body, &map_data)

	if err != nil {
		return models.MapData{}, models.Config{}, err
	}

	var config models.Config

	err = json.Unmarshal(body, &config)

	if err != nil {
		return models.MapData{}, models.Config{}, err
	}

	return map_data, config, nil
}
