package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func Fetch[T any](c *Client, url string) (T, error) {
	var zero T
	if raw_data, ok := c.cache.Get(url); ok {
		var value_data T

		err := json.Unmarshal(raw_data, &value_data)
		if err != nil {
			return zero, err
		}

		return value_data, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return zero, err
	}

	body, err := io.ReadAll(res.Body)

	c.cache.Add(url, body)

	if err != nil {
		return zero, err
	}

	res.Body.Close()

	var value_data T

	err = json.Unmarshal(body, &value_data)

	if err != nil {
		return zero, err
	}

	return value_data, nil
}
