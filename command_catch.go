package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon_name>")
	}
	name := strings.ToLower(args[0])
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	p, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return fmt.Errorf("couldn't find %s: %w", name, err)
	}

	rand.Seed(time.Now().UnixNano())
	// Simple difficulty: higher base exp => harder
	// Clamp difficulty into [10, 90]
	diff := p.BaseExperience
	if diff < 10 {
		diff = 10
	}
	if diff > 90 {
		diff = 90
	}
	roll := rand.Intn(100) // 0-99

	if roll < (100 - diff) {
		cfg.caughtPokemon[name] = p
		fmt.Printf("%s was caught!\n", name)
		fmt.Printf("You may now inspect it with the inspect command.\n")
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
