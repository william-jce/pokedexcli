package main

import (
	pokeapi "github.com/william-jce/pokedex/internal/pokeapi"
)

type config struct {
	Next          *string
	Previous      *string
	PokeapiClient pokeapi.Client
	Pokedex       map[string]pokeapi.PokemonRes
}
