package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >  ")
		if ! scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		fmt.Printf("Your command was: %s\n", words[0])
	}
}


func cleanInput(text string) []string {
	cleanedText := strings.ToLower(text)
	cleanedText = strings.TrimSpace(cleanedText)
	return strings.Fields(cleanedText)
}