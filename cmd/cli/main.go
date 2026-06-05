package main

import (
	"github.com/hansklos71/go_pokedex/infrastructure/adapters/pokeapi"
)

func main() {

	config := &AppConfig{
		limit:         20,
		offset:        -20,
		URL:           "https://pokeapi.co/api/v2/location-area/",
		PokeAPIClient: *pokeapi.NewPokeAPIClient(),
	}

	startRepl(config)

}
