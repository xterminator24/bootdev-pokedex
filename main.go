package main

import (
	"time"

	"github.com/xterminator24/bootdev-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	config := &Config{
		pokeapiClient: pokeClient,
		Pokedex:       map[string]pokeapi.Pokemon{},
	}
	startRepl(config)
}

