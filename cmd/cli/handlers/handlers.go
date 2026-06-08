package handlers

import (
	"fmt"
	"os"

	"github.com/hansklos71/go_pokedex/cmd/cli/app"
	application_services "github.com/hansklos71/go_pokedex/internal/application_services"
)

type CommandHandler struct {
	LocationService application_services.LocationService
	PokemonService  application_services.PokemonService
}

func NewCommandHandler(
	LocationService application_services.LocationService,
	PokemonService application_services.PokemonService) *CommandHandler {
	return &CommandHandler{
		LocationService: LocationService,
		PokemonService:  PokemonService,
	}
}

func (h *CommandHandler) GetCommands() map[string]app.CLICommand {
	return map[string]app.CLICommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    h.Exit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    h.Help,
		},
		"map": {
			Name:        "map",
			Description: "Displays a map of the world",
			Callback:    h.Mapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays a map of the world with pokemon",
			Callback:    h.Mapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a location to find pokemons",
			Callback:    h.Explore,
		},
		"catch": {
			Name:        "catch",
			Description: "A command to catch a pokemon.",
			Callback:    h.Catch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "A command to inspect a pokemon.",
			Callback:    h.Inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List collected pokemons",
			Callback:    h.Pokedex,
		},
	}
}

func (h *CommandHandler) Help(config *app.App, args string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range h.GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func (h *CommandHandler) Exit(config *app.App, args string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (h *CommandHandler) Mapf(config *app.App, args string) error {
	config.Next()

	if err := h.processResults(config); err != nil {
		return err
	}

	return nil
}

func (h *CommandHandler) Mapb(config *app.App, args string) error {
	config.Previous()

	if err := h.processResults(config); err != nil {
		return err
	}

	return nil
}

func (h *CommandHandler) processResults(config *app.App) error {
	loc, err := h.LocationService.ListLocations(config.Offset, config.Limit)
	if err != nil {
		return fmt.Errorf("error while getting locations data: %w", err)
	}

	for _, l := range loc {
		fmt.Println(l.Name)
	}
	return nil
}

func (h *CommandHandler) Explore(config *app.App, locationName string) error {
	fmt.Printf("Exploring %s ...\n", locationName)

	pokemons, err := h.PokemonService.ListPokemonsForLocation(locationName)
	if err != nil {
		return fmt.Errorf("error while fetching locationDetails: %w", err)
	}

	if len(pokemons) == 0 {
		fmt.Printf("No pokemon found.")
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}

func (h *CommandHandler) Catch(config *app.App, args string) error {
	pokemonName := args
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	caught, err := h.PokemonService.Catch(pokemonName)
	if err != nil {
		fmt.Printf("error while catching pokemon: %w\n", err)
		return fmt.Errorf("error while catching pokemon: %w\n", err)
	}
	if caught == false {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}

func (h *CommandHandler) Inspect(config *app.App, args string) error {
	poke, err := h.PokemonService.GetPokemonDetailsFromPokedex(args)
	if err != nil {
		return fmt.Errorf("error while fetching pokemon data: %w", err)
	}
	if poke.Name == "" {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}
	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Base XP %s\n", poke.BasedExperience)
	return nil

}

func (h *CommandHandler) Pokedex(config *app.App, args string) error {
	pokemons := h.PokemonService.ListPokemonsFromPokedex()
	fmt.Println("Your Pokedex:")
	for _, poke := range pokemons {
		fmt.Printf("- %s\n", poke.Name)
	}
	return nil
}
