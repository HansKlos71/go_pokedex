package application_services

import (
	"github.com/hansklos71/go_pokedex/internal/domain/pokemon"
	"github.com/hansklos71/go_pokedex/internal/ports"
)

type PokemonService struct {
	PokeAPIClient     ports.PokemonLocationsPort
	PokedexRepository ports.PokedexRepository
}

func NewPokemonService(PokeAPIClient ports.PokemonLocationsPort, PokedexRepository ports.PokedexRepository) *PokemonService {
	return &PokemonService{
		PokeAPIClient:     PokeAPIClient,
		PokedexRepository: PokedexRepository,
	}
}

func (s *PokemonService) GetPokemonsForLocation(locationName string) ([]pokemon.Pokemon, error) {
	pokemons, err := s.PokeAPIClient.GetPokemonsForLocation(locationName)
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (s *PokemonService) Catch(pokemonName string) (bool, error) {
	var pokemonCatched bool
	pokemon, err := s.PokeAPIClient.GetPokemonDetails(pokemonName)
	if err != nil {
		return false, err
	}

	// falculate catch chance and set pokemonCatched
	pokemonCatched = true

	if pokemonCatched == false {
		return false, nil
	}

	s.PokedexRepository.AddPokemon(pokemon.Name)
	return true, nil
}
