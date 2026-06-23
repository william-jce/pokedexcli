package main

import (
	"strings"
)

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	result := strings.Fields(words)
	return result
}
