package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xterminator24/bootdev-pokedex/internal/pokeapi"
)

type Config struct {
	pokeapiClient		pokeapi.Client
	nextLocationsURL	*string
	prevLocationsURL	*string
	Pokedex     		map[string]pokeapi.Pokemon
}

func startRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	cleanedText := strings.ToLower(text)
	cleanedText = strings.TrimSpace(cleanedText)
	words := strings.Fields(cleanedText)
	return words
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:		 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,	
		},
		"map": {
			name: 		 "map",
			description: "Get the next page of locations",
			callback:  	 commandMapf,	
		},
		"mapb": {
			name: 		 "mapb",
			description: "Get the previous page of locations",
			callback: 	 commandMapb,	
		},
		"explore": {
			name:		 "explore <location_name>",
			description: "Explore a location",
			callback:	 commandExplore,
		},
		"catch": {
			name:		 "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:		 "inspect",
			description: "Inspect the stats of a caught pokemon",
			callback:	 commandInspect,
		},
		"pokedex": {
			name:		 "pokedex",
			description: "See all pokemon you've caught",
			callback:	 commandPokedex,
		},
		"exit": {
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback:  	 commandExit,
		},
	}
}