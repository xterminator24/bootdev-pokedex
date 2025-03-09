package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location area name")
	}

	requestedArea := args[0]
	locationArea, err := config.pokeapiClient.Explore(requestedArea)
	if err != nil {
		return errors.New("area doesn't exist")
	}

	for _,encounter := range locationArea.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	
	return nil
}
