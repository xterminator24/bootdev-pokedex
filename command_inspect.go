package main

import (
	"errors"
	"fmt"

	"github.com/xterminator24/bootdev-pokedex/internal/pokeapi"
)

func commandInspect(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a single pokemon to inspect")
	}

	name := args[0]
	pokemon, exists := config.Pokedex[name]
	if !exists {
		return fmt.Errorf("%s not found in Pokedex", name)
	}

	printStats(pokemon)

	return nil
}

func printStats(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Experience: %d\n", pokemon.BaseExperience)
	fmt.Println("Stats:")
	for _,stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _,pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}
}