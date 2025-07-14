package main

import "fmt"

func commandPokedex(c *config, arg string) error {
	if len(caughtPokemon) == 0 {
		fmt.Println("You haven't caught any pokemon yet!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
