package main

import (
	pokeapi "github.com/william-jce/pokedex/internal/pokeapi"
)

type Config struct {
	Next          *string
	Previous      *string
	PokeapiClient pokeapi.Client
}
