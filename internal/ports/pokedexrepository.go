package ports

type PokedexRepository interface {
	AddPokemon(pokemonName string) (bool, error)
}
