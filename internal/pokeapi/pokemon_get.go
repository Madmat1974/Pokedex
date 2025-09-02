package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var p Pokemon
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return p, fmt.Errorf("pokemon name required")
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	// simple GET (you can add cache like your other getters later)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return p, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return p, fmt.Errorf("pokeapi error: %s", strings.TrimSpace(string(body)))
	}

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return p, err
	}
	return p, nil
}
