package application_services

import (
	"fmt"

	"github.com/hansklos71/go_pokedex/internal/domain/location"
	"github.com/hansklos71/go_pokedex/internal/ports"
)

type LocationService struct {
	PokeApiClient ports.PokemonLocationsPort
}

func NewLocationService(PokeAPIClient ports.PokemonLocationsPort) *LocationService {
	return &LocationService{
		PokeApiClient: PokeAPIClient,
	}
}

func (s *LocationService) ListLocations(offset, limit int) ([]location.Location, error) {
	locations, err := s.PokeApiClient.ListLocations(offset, limit)
	if err != nil {
		fmt.Errorf("error while fetching locations: %w", err)
		return nil, err
	}
	return locations, nil
}
