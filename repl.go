package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/time4soup/pokedex_go/internal/poke_api_client"
	"github.com/time4soup/pokedex_go/internal/pokecache"
)

// prompts input and repeatedly reads, cleans, and executes input commands matching regsitry commands
func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	pokedex := map[string]poke_api_client.Pokemon{}
	cfg := Config{
		nil,
		nil,
		pokecache.NewCache(time.Second * 5),
		[]string{},
		pokedex,
	}

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := scanner.Text()
		cfg.commands = cleanInput(input)

		if i, ok := registry()[cfg.commands[0]]; ok {
			err := i.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		fmt.Print("\nPokedex > ")
	}
}

// returns slice of individual words from raw cli input, removing all whitespace
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
