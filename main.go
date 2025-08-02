package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

type mapLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type mapLocations struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []mapLocation `json:"results"`
}

var cliCommands map[string]cliCommand
var configuration config

func init() {
	configuration = config{
		Next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Previous: "",
	}

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
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Displays the names of the last 20 location areas in the Pokemon world",
			callback:    commandMapb,
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
			err := val.callback(&configuration)
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

func commandExit(config *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for command, _ := range cliCommands {
		fmt.Printf("%v: %v\n", cliCommands[command].name, cliCommands[command].description)
	}
	return nil
}

func commandMap(config *config) error {
	if config.Next == "" {
		fmt.Println("No more locations!")
		return nil
	}

	res, err := http.Get(config.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	locationData := mapLocations{}

	err = json.Unmarshal(data, &locationData)
	if err != nil {
		return err
	}

	for _, loc := range locationData.Results {
		fmt.Println(loc.Name)
	}

	if locationData.Previous != "" {
		config.Previous = locationData.Previous
	} else {
		config.Previous = ""
	}

	if locationData.Next != "" {
		config.Next = locationData.Next
	} else {
		config.Next = ""
	}

	return nil
}

func commandMapb(config *config) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(config.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	locationData := mapLocations{}

	err = json.Unmarshal(data, &locationData)
	if err != nil {
		return err
	}

	for _, loc := range locationData.Results {
		fmt.Println(loc.Name)
	}

	if locationData.Previous != "" {
		config.Previous = locationData.Previous
	} else {
		config.Previous = ""
	}

	if locationData.Next != "" {
		config.Next = locationData.Next
	} else {
		config.Next = ""
	}

	return nil
}
