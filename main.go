package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	pokeapi "github.com/william-jce/pokedex/internal/pokeapi"
	pokecache "github.com/william-jce/pokedex/internal/pokecache"
)

var cliCommands map[string]cliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cliCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all the Pokemon in the specified location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the specified Pokemon",
			callback:    commandCatch,
		},
	}
	interval := 5 * time.Second
	cache := pokecache.NewCache(interval)
	client := pokeapi.Client{Cache: cache}
	c := config{PokeapiClient: client, Pokedex: make(map[string]pokeapi.PokemonRes)}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			continue
		}
		command, exists := cliCommands[cleanedInput[0]]
		if exists {
			var arg string
			if len(cleanedInput) > 1 {
				arg = cleanedInput[1]
			}
			err := command.callback(&c, arg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
