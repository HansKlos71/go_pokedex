package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hansklos71/go_pokedex/internal/domain/location"
	pokemon2 "github.com/hansklos71/go_pokedex/internal/domain/pokemon"
)

type PokeAPIClient struct {
	baseURL string
	URLs    map[string]string
	cache   Cache
}

func NewPokeAPIClient() *PokeAPIClient {
	return &PokeAPIClient{
		baseURL: baseURL,
		URLs: map[string]string{
			locationList:    baseURL + locationList,
			locationDetails: baseURL + locationDetails,
			pokemon:         baseURL + pokemon,
		},
		cache: *NewCache(10 * 60 * 10),
	}
}

func (c *PokeAPIClient) CallPokeAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http error while fetching location areas data: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the response body: %w", err)
	}
	return data, nil
}

func (c *PokeAPIClient) ListLocations(offset, limit int) ([]location.Location, error) {
	url := fmt.Sprintf("%s?offset=%d&limit=%d", c.URLs[locationList], offset, limit)

	data, ok := c.cache.Get(url)
	if !ok {
		var err error
		data, err = c.CallPokeAPI(url)
		if err != nil {
			return nil, fmt.Errorf("error while fetching locations data: %w", err)
		}
		c.cache.Add(url, data)
	}

	var parsedResponse LocationResponse

	if err := json.Unmarshal(data, &parsedResponse); err != nil {
		return nil, fmt.Errorf("error while parsing locations data: %w", err)
	}

	var locations []Location
	locations = parsedResponse.Results

	return c.parseLocations(locations), nil
}

func (c *PokeAPIClient) parseLocations(locations []Location) []location.Location {
	var parsedLocations []location.Location
	for _, loc := range locations {
		parsedLocations = append(parsedLocations, location.Location{Name: loc.Name})
	}
	return parsedLocations
}

func (c *PokeAPIClient) GetPokemonsForLocation(locationName string) ([]pokemon2.Pokemon, error) {
	locationDetails, err := c.getLocationDetails(locationName)
	if err != nil {
		return nil, fmt.Errorf("error while fetching location details: %w", err)
	}
	return c.parsePokemonsFromLocationDetails(locationDetails), nil

}

func (c *PokeAPIClient) getLocationDetails(locationName string) (LocationDetailsResponse, error) {
	url := fmt.Sprintf("%s/%s/", c.URLs[locationList], locationName)

	data, ok := c.cache.Get(url)

	if !ok {
		var err error
		data, err = c.CallPokeAPI(url)
		if err != nil {
			return LocationDetailsResponse{}, fmt.Errorf("error while fetching location details: %w", err)
		}
		c.cache.Add(url, data)
	}

	var locationDetailsResponse LocationDetailsResponse

	if err := json.Unmarshal(data, &locationDetailsResponse); err != nil {
		return LocationDetailsResponse{}, fmt.Errorf("error while parsing location details: %w", err)
	}

	return locationDetailsResponse, nil
}

func (c *PokeAPIClient) parsePokemonsFromLocationDetails(locationDetails LocationDetailsResponse) []pokemon2.Pokemon {
	var pokemonList []pokemon2.Pokemon
	for _, encounter := range locationDetails.PokemonEncounters {
		pokemonList = append(pokemonList, pokemon2.Pokemon{Name: encounter.Pokemon.Name})
	}
	return pokemonList
}

//func (c *PokeAPIClient) GetPokemonDetails(pokemonName string) error {
//	data, ok := c.cache.Get(pokemonName)
//	if !ok {
//		// fetch data
//	}
//
//	var parsedPokemonData
//
//	return nil
//}
