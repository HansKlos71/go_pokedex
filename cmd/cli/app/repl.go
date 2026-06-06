package app

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

func StartRepl(config *App, dispatcher Dispatcher) {
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
		cmd, ok := dispatcher.GetCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.Callback(config, optionalArg); err != nil {
			fmt.Errorf("error while executing command %s: %w\n", cmd.Name, err)
		}
	}
}
