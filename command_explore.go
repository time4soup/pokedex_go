package main

import (
	"fmt"

	"github.com/time4soup/pokedex_go/internal/poke_api_client"
)

// makes call to get info for given area and prints list of pokemon there
func commandExplore(cfg *Config) error {
	if len(cfg.commands) != 2 {
		return fmt.Errorf("invalid exploration arguments (explore <location>)")
	}
	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area/", cfg.commands[1])

	body, ok := cfg.cache.Get(url)
	var err error
	if !ok {
		body, err = poke_api_client.Get(url)
		if err != nil {
			return err
		}
		cfg.cache.Add(url, body)
	}

	locArea := UnmarshalType(body, &poke_api_client.LocationArea{})
	pokemonEncounters := locArea.PokemonEncounters

	for _, p := range pokemonEncounters {
		fmt.Printf("%s\n", p.Pokemon.Name)
	}

	return nil
}

//add ability to input arguments after explore
