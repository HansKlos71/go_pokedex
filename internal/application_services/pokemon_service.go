package application_services

import (
	"fmt"

	"github.com/hansklos71/go_pokedex/internal/domain/pokemon"
	"github.com/hansklos71/go_pokedex/internal/ports"
)

type PokemonService struct {
	PokeAPIClient ports.PokemonLocationsPort
}

func NewPokemonService(PokeAPIClient ports.PokemonLocationsPort) *PokemonService {
	return &PokemonService{
		PokeAPIClient: PokeAPIClient,
	}
}

func (s *PokemonService) GetPokemonsForLocation(locationName string) ([]pokemon.Pokemon, error) {
	pokemons, err := s.PokeAPIClient.GetPokemonsForLocation(locationName)
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (s *PokemonService) Catch(pokemonName string) error {

	return fmt.Errorf("Not implemented yet")
}
