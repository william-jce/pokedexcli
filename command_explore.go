package main

import (
	"fmt"
)

func commandExplore(c *Config, locationAreaName string) error {
	if locationAreaName == "" {
		return fmt.Errorf("Location not provided. Please enter location name.")
	}
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
