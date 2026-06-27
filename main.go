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
	}
	interval := 5 * time.Second
	cache := pokecache.NewCache(interval)
	client := pokeapi.Client{Cache: cache}
	c := Config{PokeapiClient: client}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			continue
		}
		command, exists := cliCommands[cleanedInput[0]]
		if exists {
			err := command.callback(&c)
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
