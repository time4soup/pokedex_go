package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/time4soup/pokedex_go/internal/poke_api_client"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := Config{"", ""}

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := scanner.Text()
		command := cleanInput(input)[0]

		if i, ok := registry()[command]; ok {
			i.callback(&config)
		} else {
			fmt.Println("Unknown command")
		}

		fmt.Print("\nPokedex > ")
	}
}

type Config struct {
	next     string
	previous string
}

type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func registry() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, value := range registry() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandMap(c *Config) error {
	url := c.next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	mapRes := poke_api_client.PokeApiGet(url)
	if mapRes.Next == nil {
		c.next = ""
	} else {
		c.next = *mapRes.Next
	}
	if mapRes.Previous == nil {
		c.previous = ""
	} else {
		c.previous = *mapRes.Previous
	}

	for _, item := range mapRes.Results {
		fmt.Printf("%s\n", item.Name)
	}
	return nil
}

func commandMapB(c *Config) error {
	url := c.previous
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	mapRes := poke_api_client.PokeApiGet(url)
	if mapRes.Next == nil {
		c.next = ""
	} else {
		c.next = *mapRes.Next
	}
	if mapRes.Previous == nil {
		c.previous = ""
	} else {
		c.previous = *mapRes.Previous
	}

	for _, item := range mapRes.Results {
		fmt.Printf("%s\n", item.Name)
	}
	return nil
}
