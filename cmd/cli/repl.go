package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var wierdChars = []string{",", "!", "?", ".", ";", ":"}
	for c := range wierdChars {
		text = strings.ReplaceAll(text, wierdChars[c], " ")
	}

	return strings.Fields(strings.ToLower(text))
}

func startRepl(config *AppConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		command := cleanedInput[0]
		optionalArg := ""
		if len(cleanedInput) > 1 {
			optionalArg = cleanedInput[1]
		}
		cmd, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		// fmt.Printf("Offset before cmd: %d\n", config.offset)
		if err := cmd.callback(config, optionalArg); err != nil {
			fmt.Errorf("error while executing command %s: %w\n", cmd.name, err)
		}
		// fmt.Printf("Offset after cmd: %d\n", config.offset)
	}
}
