// go
package pokeapi

import (
	"Pokedex/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	c := pokecache.NewCache(5 * time.Minute)
	return Client{
		baseURL: "https://pokeapi.co/api/v2",
		cache:   c,
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
