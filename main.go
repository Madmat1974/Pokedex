package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	var registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func() error {
			return commandHelp(registry)
		},
	}

	//s := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	var liner string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		liner = scanner.Text()
		s := cleanInput(liner)
		if len(s) == 0 {
			continue
		}
		command, exists := registry[s[0]]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commandRegistry map[string]cliCommand) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range commandRegistry {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func cleanInput(intext string) []string {
	//s := []string{}
	lowText := strings.ToLower(intext)
	trimText := strings.TrimSpace(lowText)
	s := strings.Fields(trimText)
	return s
}
