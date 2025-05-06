package main

import (
	"fmt"

	"github.com/time4soup/pokedex_go/internal/poke_api_client"
)

func commandExplore(cfg *Config) error {
	if len(cfg.commands) < 2 {
		fmt.Print(cfg.commands)
		return fmt.Errorf("missing location argument (explore <location>)")
	}
	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area/", cfg.commands[1])

	body, ok := cfg.cache.Get(url)
	if !ok {
		body = poke_api_client.Get(url)
		cfg.cache.Add(url, body)
		fmt.Printf("cached %s\n", url)
	} else {
		fmt.Printf("using cache %s\n", url)
	}

	locArea := UnmarshalType(body, &poke_api_client.LocationArea{})
	pokemonEncounters := locArea.PokemonEncounters

	for _, p := range pokemonEncounters {
		fmt.Printf("%s\n", p.Pokemon.Name)
	}

	return nil
}

//add ability to input arguments after explore
