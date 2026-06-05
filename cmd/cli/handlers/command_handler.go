package handlers

import (
	"fmt"
	"os"

	"github.com/hansklos71/go_pokedex/infrastructure/adapters/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *AppConfig, args string) error
}

type CommandHandler struct {
	Commands map[string]cliCommand
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		Commands: getCommands(),
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a map of the world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays a map of the world with pokemon",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location to find pokemons",
			callback:    commandExplore,
		},
	}
}

func (h *CommandHandler) Help(config *AppConfig, args string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func (h *CommandHandler) Exit(config *AppConfig, args string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (h *CommandHandler) Mapf(config *AppConfig, args string) error {
	config.Next()

	if err := processResults(config); err != nil {
		return err
	}

	return nil
}

func (h *CommandHandler) Mapb(config *AppConfig, args string) error {
	config.Previous()

	if err := processResults(config); err != nil {
		return err
	}

	return nil
}

func (h *CommandHandler) processResults(config *AppConfig) error {
	loc, err := config.PokeAPIClient.GetLocations(config.offset, config.limit)
	if err != nil {
		fmt.Errorf("error while getting locations data: %w", err)
		return err
	}

	for _, l := range loc {
		fmt.Println(l.Name)
	}

	return nil
}

func (h *CommandHandler) Explore(config *AppConfig, locationName string) error {
	fmt.Printf("Exploring %s ...\n", locationName)
	locationDetails, err := config.PokeAPIClient.GetLocationDetails(locationName)
	if err != nil {
		return fmt.Errorf("error while fetching locationDetails: %w", err)
	}

	pokemons := pokeapi.ParsePokemonsFromLocationDetails(locationDetails)
	if len(pokemons) == 0 {
		fmt.Printf("No pokemon found.")
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons {
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}

func (h *CommandHandler) Catch(config *AppConfig, args string) error {
	pokemonName := args
	fmt.Printf("Throwing a Pokeball at %s...", pokemonName)

	return nil
}
