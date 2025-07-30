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
		fmt.Print("Pokédex >")
		scanner.Scan()
		input := scanner.Text()
		cleaned_input := cleanInput(input)
		first_word := cleaned_input[0]
		fmt.Printf("Your command was: %v\n", first_word)
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	textSlice := strings.Fields(lowercaseText)
	return textSlice
}
