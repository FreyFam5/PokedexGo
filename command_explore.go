package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/FreyFam5/go/pokedexcli/internal/pokeapi"
)

func commandExplore(c *config, arg string) error {
	if arg == "" {
		return fmt.Errorf("error: needs an area name / id")
	}
	fmt.Printf("Exploring %s...\n", arg)
	fullUrl := mapPrefix + arg
	body, err := getBodyWithCache(fullUrl)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			fmt.Printf("'%s' not found!\n", arg)
			return nil
		}
		return err
	}

	areaInfo := pokeapi.AreaInfo{}
	if err := json.Unmarshal(body, &areaInfo); err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for idx := range areaInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", areaInfo.PokemonEncounters[idx].Pokemon.Name)
	}
	return nil
}
