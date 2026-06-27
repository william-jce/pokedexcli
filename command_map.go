package main

import (
	"fmt"
)

func commandMap(c *Config, locationAreaName string) error {
	locations, err := c.PokeapiClient.GetLocationArea(c.Next)
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
