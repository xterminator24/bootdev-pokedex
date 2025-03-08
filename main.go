package main

import (
	"time"

	"github.com/xterminator24/bootdev-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(120 * time.Second)
	config := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}

