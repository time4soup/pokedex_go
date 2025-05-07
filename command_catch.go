package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/time4soup/pokedex_go/internal/poke_api_client"
)

// gets data for given pokemon and the randomly determines if caught, adds pokemon to pokedex if caught
func commandCatch(cfg *Config) error {
	if len(cfg.commands) != 2 {
		return fmt.Errorf("invalid catch arguments (catch <pokemon>)")
	}
	url := fmt.Sprint("https://pokeapi.co/api/v2/pokemon/", cfg.commands[1])

	body, ok := cfg.cache.Get(url)
	var err error
	if !ok {
		body, err = poke_api_client.Get(url)
		if err != nil {
			return err
		}
		cfg.cache.Add(url, body)
	}

	pokemon := UnmarshalType(body, &poke_api_client.Pokemon{})
	pokeSpecies, err := getSpecies(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	isCaught := rand.Intn(256) < pokeSpecies.CaptureRate //255 is max capture rate
	switch isCaught {
	case true:
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = *pokemon
	case false:
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

// makes call to get pokemon species info to be used for catch rate
func getSpecies(cfg *Config) (*poke_api_client.PokemonSpecies, error) {
	url := fmt.Sprint("https://pokeapi.co/api/v2/pokemon-species/", cfg.commands[1])

	body, ok := cfg.cache.Get(url)
	var err error
	if !ok {
		body, err = poke_api_client.Get(url)
		if err != nil {
			return nil, err
		}
		cfg.cache.Add(url, body)
	}

	return UnmarshalType(body, &poke_api_client.PokemonSpecies{}), nil
}
