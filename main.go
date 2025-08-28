package main

import (
	"Pokedex/internal/pokeapi"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type commandFn func(cfg *Config, args []string, registry map[string]cliCommand) error

type cliCommand struct {
	name        string
	description string
	callback    commandFn
}

type Config struct {
	Next     *string
	Previous *string
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

func commandHelp(cfg *Config, args []string, registry map[string]cliCommand) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range registry {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap(cfg *Config, args []string, registry map[string]cliCommand) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil || cfg.Previous != nil {
		if cfg.Next != nil {
			url = *cfg.Next
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data pokeapi.LocationAreaList

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	for _, r := range data.Results {
		fmt.Println(r.Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous
	return nil

}

func commandMapb(cfg *Config, args []string, registry map[string]cliCommand) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.Previous

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data pokeapi.LocationAreaList
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	for _, r := range data.Results {
		fmt.Println(r.Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous
	return nil

}

func commandExit(cfg *Config, args []string, registry map[string]cliCommand) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(intext string) []string {
	//s := []string{}
	lowText := strings.ToLower(intext)
	trimText := strings.TrimSpace(lowText)
	s := strings.Fields(trimText)
	return s
}
