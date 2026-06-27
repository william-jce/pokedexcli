package main

import (
	"strings"
)

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	result := strings.Fields(words)
	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
