package app

import (
	"github.com/hansklos71/go_pokedex/infrastructure/adapters/pokeapi"
	"github.com/hansklos71/go_pokedex/internal/application_services"
)

type CLICommand struct {
	Name        string
	Description string
	Callback    func(config *App, args string) error
}

type Dispatcher interface {
	GetCommands() map[string]CLICommand
}

type App struct {
	Limit           int
	Offset          int
	URL             string
	PokeAPIClient   pokeapi.PokeAPIClient
	LocationService application_services.LocationService
	PokemonService  application_services.PokemonService
}

func (c *App) Next() {
	c.Offset += 20
	return
}

func (c *App) Previous() {
	if c.Offset > 20 {
		c.Offset -= 20
		return
	}
	c.Offset = 0
}
