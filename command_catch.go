package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon to catch")
	}

	name := args[0]
	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// Generate a random number between 0 and 100
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	randomInt := r.Intn(pokemon.BaseExperience)
	percentageForCatch := float32(0.72)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if float32(randomInt) <= float32(pokemon.BaseExperience) * percentageForCatch {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	// Pokemon caught add to pokedex
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You mayu now inspect it with the inspect command")
	config.Pokedex[pokemon.Name] = pokemon

	return nil
}

