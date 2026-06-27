package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Location not provided. Please enter location name.")
	}
	locationAreaName := args[0]
	location, err := c.PokeapiClient.GetAreaPokemon(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationAreaName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
