package application_services

import (
	"fmt"
	"math/rand"

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
	if rand.Intn(200) > pokemon.BasedExperience {
		pokemonCatched = false
		return pokemonCatched, nil
	}

	pokemonCatched = true
	added, err := s.PokedexRepository.AddPokemon(pokemon.Name)
	if err != nil {
		fmt.Errorf("error adding pokemon to pokedex: %w", err)
	} else if added != true {
		fmt.Println("already in pokedex")
	}
	return pokemonCatched, nil
}

func (s *PokemonService) GetPokemonsDataFromPokedex(pokemonName string) (pokemon.Pokemon, error) {
	poke, err := s.PokedexRepository.GetPokemonByName(pokemonName)
	if err != nil {
		return pokemon.Pokemon{}, err
	}

	return poke, nil
}

func (s *PokemonService) ListPokemonsFromPokedex() []pokemon.Pokemon {
	return s.PokedexRepository.ListPokemons()
}
