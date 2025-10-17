package pokeapi

import (
	"net/http"

	"github.com/sshaheen/pokedexcli/internal/pokecache"
)

type Client struct {
	cache  *pokecache.Cache
	client http.Client
}

func NewClient(c *pokecache.Cache) *Client {
	return &Client{cache: c, client: http.Client{}}
}
