package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
var config Config

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&config)
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

type Config struct {
	Next	string
	Prev	string
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:		 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,	
		},
		"exit": {
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback:  	 commandExit,
		},
		"map": {
			name: 		 "map",
			description: "Fetch the next 20 map location areas",
			callback:  	 commandMap,	
		},
		"mapb": {
			name: 		 "mapb",
			description: "Fetch the previous 20 map location areas",
			callback: 	 commandMapb,	
		},
	}
}