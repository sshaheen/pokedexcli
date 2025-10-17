package models

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
