package main

import (
	"fmt"

	"github.com/time4soup/pokedex_go/internal/poke_api_client"
)

// makes get call to location-area in pokeapi and prints previous 20 locations
func commandMapB(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	if cfg.previous != nil {
		url = *cfg.previous
	}

	body, ok := cfg.cache.Get(url)
	var err error
	if !ok {
		body, err = poke_api_client.Get(url)
		if err != nil {
			return err
		}
		cfg.cache.Add(url, body)
	}

	locList := UnmarshalType(body, &poke_api_client.LocationAreaList{})

	cfg.next = locList.Next
	cfg.previous = locList.Previous

	for _, item := range locList.Results {
		fmt.Printf("%s\n", item.Name)
	}
	return nil
}
