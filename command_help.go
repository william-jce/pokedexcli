package main

import (
	"fmt"
)

func commandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for key, command := range cliCommands {
		fmt.Printf("%s: %s\n", key, command.description)
	}
	return nil
}
