package application_services

import (
	"fmt"

	"github.com/hansklos71/go_pokedex/internal/ports"
)

type Pokemon struct {
	Name string
}
 a
type PokemonService struct {
	Locationclient ports.PokemonLocationsPort
}

func NewPokemonService() *PokemonService {
	return &PokemonService{}
}

func (s *PokemonService) GetPokemonsForLocation(locationName string) ([]Pokemon, error) {
	pokemons, err := s.GetPokemonsForLocation(locationName)
	if err != nil {
		fmt.Errorf("error while fetching location details: %w", err)
		return nil, err
	}
	return pokemons, nil
}

func (s *PokemonService) Catch(pokemonName string) error {
	fmt.Println("Not implemented.")
	return nil
}
