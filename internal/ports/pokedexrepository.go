package ports

import "github.com/hansklos71/go_pokedex/internal/domain/pokemon"

type PokedexRepository interface {
	AddPokemon(pokemonName string) (bool, error)
	GetPokemonByName(pokemonName string) (pokemon.Pokemon, error)
	ListPokemons() []pokemon.Pokemon
}
