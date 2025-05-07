package main

import (
	"fmt"
)

// prints basic stats of pokemon if in the pokedex
func commandInspect(cfg *Config) error {
	if len(cfg.commands) != 2 {
		return fmt.Errorf("invalid inspect arguments (inpect <pokemon>)")
	}
	pokemon, ok := cfg.pokedex[cfg.commands[1]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("   -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("   - %s\n", pokeType.Type.Name)
	}

	return nil
}
