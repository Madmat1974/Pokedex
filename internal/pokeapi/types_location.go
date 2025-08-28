package pokeapi

// go

type NamedAPIRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaList struct {
	Next     *string       `json:"next"`
	Previous *string       `json:"previous"`
	Results  []NamedAPIRef `json:"results"`
}
