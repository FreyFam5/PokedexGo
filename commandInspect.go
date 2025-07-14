package main

import "fmt"

func commandInspect(c *config, arg string) error {
	pokemon, exists := caughtPokemon[arg]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, mon := range pokemon.Types {
		fmt.Printf("  - %s\n", mon.Type.Name)
	}
	return nil
}
