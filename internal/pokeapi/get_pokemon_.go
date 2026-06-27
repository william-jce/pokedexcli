package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetPokemon(pokemonName string) (PokemonRes, error) {
	var pokemon PokemonRes
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	fullURL := baseURL + pokemonName

	if entry, ok := c.Cache.Get(fullURL); ok {
		if err := json.Unmarshal(entry, &pokemon); err != nil {
			return PokemonRes{}, err
		}
		return pokemon, nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return PokemonRes{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonRes{}, err
	}
	if res.StatusCode > 299 {
		return PokemonRes{}, fmt.Errorf("response failed with status code %d", res.StatusCode)
	}
	c.Cache.Add(fullURL, data)

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return PokemonRes{}, err
	}

	return pokemon, nil
}
