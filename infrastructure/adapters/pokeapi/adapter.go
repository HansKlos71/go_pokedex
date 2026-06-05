package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	application_services "github.com/hansklos71/go_pokedex/internal/application services"
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

func (c *PokeAPIClient) GetLocations(offset, limit int) ([]Location, error) {
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
	return locations, nil
}

func (c *PokeAPIClient) GetLocationDetails(locationName string) (LocationDetailsResponse, error) {
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

	var parsedLocationDetailsResponse LocationDetailsResponse

	if err := json.Unmarshal(data, &parsedLocationDetailsResponse); err != nil {
		return LocationDetailsResponse{}, fmt.Errorf("error while parsing location details: %w", err)
	}

	return parsedLocationDetailsResponse, nil
}

func (c *PokeAPIClient) GetPokemonForLocation(locationName string) ([]application_services.Pokemon, error) {
	locationDetails, err := c.GetLocationDetails(locationName)
	if err != nil {
		return nil, fmt.Errorf("error while fetching location details: %w", err)
	}
	return c.parsePokemonsFromLocationDetails(locationDetails), nil

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

func (c *PokeAPIClient) parsePokemonsFromLocationDetails(locationDetails LocationDetailsResponse) []application_services.Pokemon {
	var pokemonList []application_services.Pokemon
	for _, encounter := range locationDetails.PokemonEncounters {
		pokemonList = append(pokemonList, application_services.Pokemon{Name: encounter.Pokemon.Name})
	}
	return pokemonList
}
