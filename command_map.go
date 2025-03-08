package main

import (
	"errors"
	"fmt"

	"github.com/xterminator24/bootdev-pokedex/internal/pokeapi"
)

func commandMapf(config *Config) error {

	locationsResp, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationsResp.Next
	config.prevLocationsURL = locationsResp.Previous

	printLocationAreas(locationsResp)

	return nil
}

func commandMapb(config *Config) error {
	if config.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationsResp.Next
	config.prevLocationsURL = locationsResp.Previous

	printLocationAreas(locationsResp)

	return nil
}

func printLocationAreas(locationsResp pokeapi.RespShallowLocations) {
	for _,location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
}