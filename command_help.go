package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for key, command := range cliCommands {
		fmt.Printf("%s: %s\n", key, command.description)
	}
	return nil
}
