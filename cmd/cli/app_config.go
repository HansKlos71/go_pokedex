package main

import (
	"github.com/hansklos71/go_pokedex/infrastructure/adapters/pokeapi"
)

type AppConfig struct {
	limit         int
	offset        int
	URL           string
	PokeAPIClient pokeapi.PokeAPIClient
}

func (c *AppConfig) Next() {
	c.offset += 20
	return
}

func (c *AppConfig) Previous() {
	if c.offset > 20 {
		c.offset -= 20
		return
	}
	c.offset = 0
}
