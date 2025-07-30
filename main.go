package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommands map[string]cliCommand

func init() {
	cliCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokédex >")
		scanner.Scan()
		input := scanner.Text()
		cleaned_input := cleanInput(input)

		if len(cleaned_input) == 0 {
			fmt.Print("Please enter a command\n")
			continue
		}
		command := cleaned_input[0]

		val, ok := cliCommands[command]
		if ok {
			err := val.callback()
			if err != nil {
				fmt.Printf("Error occured: %v", err)
			}
		} else {
			fmt.Print("Unkown command\n")
		}

	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	textSlice := strings.Fields(lowercaseText)
	return textSlice
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for command, _ := range cliCommands {
		fmt.Printf("%v: %v\n", cliCommands[command].name, cliCommands[command].description)
	}
	return nil
}
