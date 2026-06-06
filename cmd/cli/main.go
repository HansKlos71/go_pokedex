package main

import (
	config2 "github.com/hansklos71/go_pokedex/cmd/cli/app"
	"github.com/hansklos71/go_pokedex/cmd/cli/handlers"
	"github.com/hansklos71/go_pokedex/infrastructure/adapters/pokeapi"
	"github.com/hansklos71/go_pokedex/infrastructure/adapters/repository"
	"github.com/hansklos71/go_pokedex/internal/application_services"
)

func main() {

	pokedexRepository := repository.NewMemoryPokedexRepository()
	pokeAPIClient := pokeapi.NewPokeAPIClient()
	locationService := application_services.NewLocationService(pokeAPIClient)
	pokemonService := application_services.NewPokemonService(pokeAPIClient, pokedexRepository)
	commandHandler := handlers.NewCommandHandler(*locationService, *pokemonService)

	config := &config2.App{
		Limit:  20,
		Offset: 0,
	}

	config2.StartRepl(config, commandHandler)

}
