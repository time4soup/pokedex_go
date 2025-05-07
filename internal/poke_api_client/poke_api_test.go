package poke_api_client

import (
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		url string
	}{
		{
			url: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		},
		{
			url: "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
		},
	}

	for _, c := range cases {
		dat, _ := Get(c.url)

		if dat == nil {
			t.Errorf("data not retreived. url: %s", c.url)
		}
	}
}
