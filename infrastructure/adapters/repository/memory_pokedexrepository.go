package repository

import (
	"fmt"

	"github.com/hansklos71/go_pokedex/internal/domain/pokemon"
)

type MemoryPokedexRepository struct {
	pokemons map[string]pokemon.Pokemon
}

func (r *MemoryPokedexRepository) ListPokemons() []pokemon.Pokemon {
	pokeList := make([]pokemon.Pokemon, 0)
	for _, poke := range r.pokemons {
		pokeList = append(pokeList, poke)
	}
	return pokeList
}

func (r *MemoryPokedexRepository) GetPokemonByName(pokemonName string) (pokemon.Pokemon, error) {
	if _, ok := r.pokemons[pokemonName]; !ok {
		return pokemon.Pokemon{}, fmt.Errorf("pokemon not found")
	}
	return r.pokemons[pokemonName], nil
}

func NewMemoryPokedexRepository() *MemoryPokedexRepository {
	return &MemoryPokedexRepository{
		pokemons: make(map[string]pokemon.Pokemon, 0),
	}
}

func (r *MemoryPokedexRepository) GetPokemons() []string {
	var pokemons []string
	for _, poke := range r.pokemons {
		pokemons = append(pokemons, poke.Name)
	}
	return pokemons
}

func (r *MemoryPokedexRepository) AddPokemon(pokemonName string) (bool, error) {
	if _, ok := r.pokemons[pokemonName]; ok {
		return false, nil
	}
	r.pokemons[pokemonName] = pokemon.Pokemon{Name: pokemonName}
	return true, nil
}
