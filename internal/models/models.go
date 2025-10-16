package models

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

type Config struct {
	Next     string
	Previous string
}

type MapData struct {
	Results []Location
}

type Location struct {
	Name string
	Url  string
}
