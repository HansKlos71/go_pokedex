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

func (p Pokemon) GetName() string {
	return p.Name
}

func (p Pokemon) GetBasedExperience() int {
	return p.BasedExperience
}
