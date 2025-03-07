package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

var commands map[string]cliCommand

func main() {
	commands = map[string]cliCommand{
	"exit": {
		name:		 "exit",
		description: "Exit the Pokedex",
		callback:	 commandExit,
	},
	"help": {
		name:		 "help",
		description: "Displays a help message",
		callback:	 commandHelp,
	},
}
	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >  ")
		if ! scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		
		if cmd, ok := commands[words[0]]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}


func cleanInput(text string) []string {
	cleanedText := strings.ToLower(text)
	cleanedText = strings.TrimSpace(cleanedText)
	return strings.Fields(cleanedText)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _,cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

