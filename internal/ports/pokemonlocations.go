package ports

import (
	"github.com/hansklos71/go_pokedex/internal/domain/location"
	"github.com/hansklos71/go_pokedex/internal/domain/pokemon"
)

type PokemonLocationsPort interface {
	ListLocations(offset, limit int) ([]location.Location, error)
	GetPokemonsForLocation(locationName string) ([]pokemon.Pokemon, error)
	GetPokemonDetails(pokemonName string) (pokemon.Pokemon, error)
}
