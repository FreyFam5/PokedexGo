package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/FreyFam5/go/pokedexcli/internal/pokeapi"
)

const pokemonPrefix = "https://pokeapi.co/api/v2/pokemon/"

var caughtPokemon = make(map[string]pokeapi.Pokemon)

func commandCatch(c *config, arg string) error {
	if arg == "" {
		return fmt.Errorf("error: needs a pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	fullUrl := pokemonPrefix + arg
	body, err := getBodyWithCache(fullUrl)
	if err != nil || arg == "" {
		if strings.Contains(err.Error(), "404") {
			fmt.Printf("'%s' not found!\n", arg)
			return nil
		}
		return err
	}

	pokemon := pokeapi.Pokemon{}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	chance := (256.0 - float64(pokemon.BaseExperience)) / 255 // 255 is the max amount of base experience a mon can have, I subtract by 256 so theres always a chance to catch the mon
	if chance >= rand.Float64() {
		fmt.Printf("%s was caught!\n", arg)
		caughtPokemon[pokemon.Name] = pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", arg)
	return nil
}
