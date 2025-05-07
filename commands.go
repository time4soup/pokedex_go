package main

import (
	"encoding/json"
	"log"

	"github.com/time4soup/pokedex_go/internal/poke_api_client"
	"github.com/time4soup/pokedex_go/internal/pokecache"
)

// stores info to be passed to and from functions
type Config struct {
	next     *string
	previous *string
	cache    *pokecache.Cache
	commands []string
	pokedex  map[string]poke_api_client.Pokemon
}

// stores information used to interface cli with command functions
type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

// returns map with all of the CliCommand structs for each command to be used in 'help' command
func registry() map[string]CliCommand {
	return map[string]CliCommand{
		"catch": {
			name:        "catch",
			description: "Tries to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explores an area to find which Pokemon are there",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects the info about Pokemon you have caught",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Lists map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous page of map locations",
			callback:    commandMapB,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all of the Pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

// unmarshalls json data for type given by jsonType
func UnmarshalType[T any](body []byte, jsonType *T) *T {
	var data T
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	return &data
}
