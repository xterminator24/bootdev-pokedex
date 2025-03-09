package main

import (
	"fmt"
)

func commandPokedex(config *Config, args ...string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for _,pokemon := range config.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}