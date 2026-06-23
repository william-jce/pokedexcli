package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var result []string

	for word := range strings.FieldsSeq(text) {
		result = append(result, strings.ToLower(word))
	}

	return result
}
