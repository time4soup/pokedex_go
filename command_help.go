package main

import (
	"fmt"
)

// prints name, and description for each command using registry info
func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, value := range registry() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
