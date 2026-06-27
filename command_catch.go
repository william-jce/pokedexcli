package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Pokemon name not entered. Please enter Pokemon name.")
	}
	pokemonName := args[0]
	pokemon, err := c.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	difficulty := pokemon.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if rand.Intn(difficulty+50) < difficulty {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		fmt.Printf("%s was caught!\n", pokemonName)
		c.Pokedex[pokemonName] = pokemon
	}

	return nil
}
