package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location area name")
	}

	requestedArea := args[0]
	locationArea, err := config.pokeapiClient.GetLocation(requestedArea)
	if err != nil {
		return errors.New("area doesn't exist")
	}

	fmt.Printf("Exploring %s...\n", locationArea.Name)
	fmt.Printf("Found Pokemon: \n")

	for _,encounter := range locationArea.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	
	return nil
}
