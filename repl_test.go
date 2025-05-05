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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " CHEEZE GROMIT",
			expected: []string{"cheeze", "gromit"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("length mismatch for test: %s. Expected: %d, Actual %d", c.input, len(c.expected), len(actual))
		}
		for i, word := range actual {
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("unmatched words at %d. Expected: %s, Actual: %s", i, expectedWord, word)
			}
		}
	}
}
