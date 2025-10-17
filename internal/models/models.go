package models

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type MapData struct {
	Results  []Location `json:"results"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
}

type AreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
