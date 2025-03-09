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
		"mapf": {
			name: 		 "mapf",
			description: "Fetch the next 20 map location areas",
			callback:  	 commandMapf,	
		},
		"mapb": {
			name: 		 "mapb",
			description: "Fetch the previous 20 map location areas",
			callback: 	 commandMapb,	
		},
		"explore": {
			name:		 "explore",
			description: "Show pokemon in the provided area",
			callback:	 commandExplore,
		},
		"catch": {
			name:		 "catch",
			description: "Attempt to catch the provided pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback:  	 commandExit,
		},
	}
}