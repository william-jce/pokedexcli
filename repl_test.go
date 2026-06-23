package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " thisISApokemon but     no",
			expected: []string{"thisisapokemon", "but", "no"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expectedSlice := c.expected
		if len(actual) != len(expectedSlice) {
			t.Fatalf("expected: %v, got : %v", len(expectedSlice), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Fatalf("expected: %v, got: %v", expectedWord, word)
			}
		}
	}
}
