package repository

type MemoryPokedexRepository struct {
	pokemons []string
}

func NewMemoryPokedexRepository() *MemoryPokedexRepository {
	return &MemoryPokedexRepository{
		pokemons: make([]string, 0),
	}
}

func (r *MemoryPokedexRepository) GetPokemons() []string {
	return r.pokemons
}

func (r *MemoryPokedexRepository) AddPokemon(pokemonName string) (bool, error) {
	r.pokemons = append(r.pokemons, pokemonName)
	return true, nil
}
