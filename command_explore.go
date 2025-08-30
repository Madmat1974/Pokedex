package main

import (
	"context"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		fmt.Println("usage: explore <area_name>")
		return nil
	}
	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	ctx := context.Background()
	areaResp, err := cfg.pokeapiClient.GetLocationArea(ctx, areaName)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, e := range areaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", e.Pokemon.Name)
	}
	return nil
}
