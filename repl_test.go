package main

import (
	"testing"
	"time"

	"github.com/time4soup/pokedex_go/internal/pokecache"
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

func TestCaching(t *testing.T) {
	cases := []struct {
		url    string
		cached bool
		wait   time.Duration
	}{
		{
			url:    "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
			cached: true,
			wait:   0,
		},
		{
			url:    "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
			cached: false,
			wait:   0,
		},
		{
			url:    "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
			cached: false,
			wait:   time.Second * 3,
		},
	}
	config := Config{
		nil,
		nil,
		pokecache.NewCache(time.Second),
		[]string{},
	}

	body, cached := config.cache.Get("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")
	if !cached {
		config.cache.Add("https://pokeapi.co/api/v2/location-area?offset=0&limit=20", body)
	}

	for _, c := range cases {
		time.Sleep(c.wait)
		body, cached := config.cache.Get(c.url)
		if !cached {
			config.cache.Add(c.url, body)
		}

		if cached != c.cached {
			t.Errorf("incorrect cache. Expected: %v. Actual: %v. url: %s", c.cached, cached, c.url)
		}
	}
}
