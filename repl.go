package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ") // The starting line of each input

		scanner.Scan() // Pushes the scan to await next input

		input := scanner.Text() // The text found in the scan
		input = cleanInput(input)[0] // Cleans the input

		command, exists := commands[input] // Finds the command if it exists
		if exists {
			if err := command.callback(&Config{}); err != nil{ // If it doesn't exist, will print error and skip to next input
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
	Name 	string
	Url 	string
}

type Config struct {
	Count 		int
	Next 			string
	Previous 	string
	Results 	[]NamedAPIResource
}

type cliCommand struct {
	name 				string
	description string
	callback 		func(*Config) error
}

var commands = map[string]cliCommand{}

// Sets up commands map after initializing to avoid a circular initialization with commands that use this
func init() {
	commands = map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of 20 location ares in Pokemon, each subsequent call will display the next 20 and so forth",
			callback: commandMap,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}
