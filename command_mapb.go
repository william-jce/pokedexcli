package main

import (
	"fmt"
)

func commandMapb(c *Config, locationAreaName string) error {
	if c.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	locations, err := c.PokeapiClient.GetLocationArea(c.Previous)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
