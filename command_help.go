package main

import (
	"fmt"
)

func commandHelp(c *config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, value := range commands {
		fmt.Println(value.name + ": " + value.description)
	}
	return nil
}
