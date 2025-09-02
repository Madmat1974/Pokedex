package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon_name>")
	}
	name := args[0]

	val, ok := cfg.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Println("Name:", val.Name)
	fmt.Println("Height:", val.Height)
	fmt.Println("Weight:", val.Weight)
	fmt.Println("Stats:")
	for _, s := range val.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range val.Types {
		fmt.Println("  -", t.Type.Name)
	}
	return nil

}
