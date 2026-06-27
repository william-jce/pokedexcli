package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocationArea(url *string) (LocationAreaRes, error) {
	var locationArea LocationAreaRes
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if url != nil {
		baseURL = *url
	}

	if entry, ok := c.Cache.Get(baseURL); ok {
		if err := json.Unmarshal(entry, &locationArea); err != nil {
			return LocationAreaRes{}, err
		}
		return locationArea, nil
	}

	res, err := http.Get(baseURL)
	if err != nil {
		return LocationAreaRes{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}
	if res.StatusCode > 299 {
		return LocationAreaRes{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	c.Cache.Add(baseURL, data)

	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationAreaRes{}, err
	}

	return locationArea, nil
}
