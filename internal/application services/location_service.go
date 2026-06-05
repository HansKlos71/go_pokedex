package application_services

import (
	"fmt"

	"github.com/hansklos71/go_pokedex/internal/ports"
)

type Location struct {
	Name string
}

type LocationService struct {
	Client ports.PokemonLocationsPort
}

func NewLocationService(client ports.PokemonLocationsPort) *LocationService {
	return &LocationService{
		Client: client,
	}
}

func (s *LocationService) ListLocations(offset, limit int) ([]Location, error) {
	locations, err := s.Client.ListLocations(offset, limit)
	if err != nil {
		fmt.Errorf("error while fetching locations: %w", err)
		return nil, err
	}
	return locations, nil
}
