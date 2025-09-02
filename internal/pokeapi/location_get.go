package pokeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// go
func (c *Client) GetLocationArea(ctx context.Context, name string) (LocationArea, error) {
	url := c.baseURL + "/location-area/" + name

	if b, ok := c.cache.Get(url); ok {
		var area LocationArea
		if err := json.Unmarshal(b, &area); err != nil {
			return LocationArea{}, err
		}
		return area, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, b)

	var area LocationArea
	if err := json.Unmarshal(b, &area); err != nil {
		return LocationArea{}, err
	}
	return area, nil

}
