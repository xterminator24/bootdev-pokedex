package main

import (
	"errors"
	"fmt"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon to catch")
	}

	foundPokemon := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", foundPokemon)

	caughtPokemon, err := config.pokeapiClient.CatchPokemon(foundPokemon)
	if err != nil {
		fmt.Printf("%s escaped!", foundPokemon)
		return err
	}

	// Pokemon caught add to pokedex
	fmt.Printf("%s was caught!\n", foundPokemon)
	config.pokeapiClient.Pokedex[foundPokemon] = caughtPokemon

	return nil
}

