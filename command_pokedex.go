package main

import (
	"fmt"
)

func commandPokedex(c *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemonName := range c.Pokedex {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
