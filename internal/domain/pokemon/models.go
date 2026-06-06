package pokemon

type Pokemon struct {
	Name            string
	BasedExperience int
}

func NewPokemon(name string, basedExperience *int) Pokemon {
	if basedExperience != nil {
		return Pokemon{
			Name:            name,
			BasedExperience: *basedExperience,
		}
	}
	return Pokemon{
		Name: name,
	}
}
