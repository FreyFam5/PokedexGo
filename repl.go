package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	presetConfig := config{} // A preset config so the commands can run with a clean config at the beginning
	for {
		fmt.Print("Pokedex > ") // The starting line of each input

		scanner.Scan() // Pushes the scan to await next input

		input := scanner.Text()     // The text found in the scan
		inputs := cleanInput(input) // Cleans the input
		input = inputs[0]
		arg := ""
		if len(inputs) > 1 {
			arg = inputs[1]
		}
		command, exists := commands[input] // Finds the command if it exists
		if exists {
			if err := command.callback(&presetConfig, arg); err != nil { // If it doesn't exist, will print error and skip to next input
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
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

type NamedAPIResource struct {
	Name string
	Url  string
}

type config struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

var commands = map[string]cliCommand{}

// Sets up commands map after initializing to avoid a circular initialization with commands that use this
func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 areas in Pokemon, each subsequent call will display the next 20 and so forth",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the 20 previous areas in Pokemon, each subsequent call will display the next 20 until reaching the first 20, in which it will stop",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Shows the pokemon that are in the given area or area id, ex. Pokedex > explore pastoria-city-area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the given pokemon, ex. Pokedex > catch pikachu",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a cuaght pokemons stats, ex. Pokedex > inspect pikachu",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows all the pokemon you've caught so far",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
