package main

import (
	"fmt"
)

// prints list of pokemon in pokedex if any are caught
func commandPokedex(cfg *Config) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caught any Pokemon yet")
	}
	fmt.Println("Your Pokedex:")
	for key := range cfg.pokedex {
		fmt.Printf("   -%s\n", key)
	}

	return nil
}
