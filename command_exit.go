package main

import (
	"fmt"
	"os"
)

// prints exit message and closes poke cli
func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
