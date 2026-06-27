package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetAreaPokemon(locationAreaName string) (AreaPokemonRes, error) {
	var areaPokemon AreaPokemonRes
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	fullURL := baseURL + locationAreaName

	if entry, ok := c.Cache.Get(fullURL); ok {
		if err := json.Unmarshal(entry, &areaPokemon); err != nil {
			return AreaPokemonRes{}, err
		}
		return areaPokemon, nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return AreaPokemonRes{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaPokemonRes{}, err
	}
	if res.StatusCode > 299 {
		return AreaPokemonRes{}, fmt.Errorf("response failed with status code %d", res.StatusCode)
	}
	c.Cache.Add(fullURL, data)

	if err := json.Unmarshal(data, &areaPokemon); err != nil {
		return AreaPokemonRes{}, err
	}

	return areaPokemon, nil
}
