package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocationArea(url *string) (LocationAreaRes, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if url != nil {
		baseURL = *url
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

	var locationArea LocationAreaRes
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationAreaRes{}, err
	}

	return locationArea, nil
}
