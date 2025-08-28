package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// go
type Config struct {
	Next     *string
	Previous *string
}

type commandFn func(cfg *Config, args []string, registry map[string]cliCommand) error

type cliCommand struct {
	name        string
	description string
	callback    commandFn
}

func main() {
	cfg := &Config{}

	registry := map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
	}
	registry["help"] = cliCommand{
		name: "help", description: "Displays a help message", callback: commandHelp,
	}

	registry["map"] = cliCommand{name: "map", description: "List next 20 location areas", callback: commandMap}
	registry["mapb"] = cliCommand{name: "mapb", description: "List previous 20 location areas", callback: commandMapb}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			return
		}
		line := scanner.Text()
		parts := cleanInput(line)
		if len(parts) == 0 {
			continue
		}
		cmdName := parts[0]
		args := parts[1:]

		cmd, ok := registry[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg, args, registry); err != nil {
			fmt.Println(err)
		}
	}
}

func commandExit(cfg *Config, args []string, registry map[string]cliCommand) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args []string, registry map[string]cliCommand) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range registry {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap(cfg *Config, args []string, registry map[string]cliCommand) error {
	return nil
}

func commandMapb(cfg *Config, args []string, registry map[string]cliCommand) error {
	return nil
}

func cleanInput(intext string) []string {
	//s := []string{}
	lowText := strings.ToLower(intext)
	trimText := strings.TrimSpace(lowText)
	s := strings.Fields(trimText)
	return s
}
