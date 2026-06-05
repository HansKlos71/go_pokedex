package ports

import application_services "github.com/hansklos71/go_pokedex/internal/application services"

type PokemonLocationsPort interface {
	ListLocations(offset, limit int) ([]application_services.Location, error)
	GetPokemonsForLocation(locationName string) (application_services.Pokemon, error)
}
