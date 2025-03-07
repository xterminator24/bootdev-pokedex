package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}


func cleanInput(text string) []string {
	cleanedText := strings.ToLower(text)
	cleanedText = strings.TrimSpace(cleanedText)
	return strings.Fields(cleanedText)
}