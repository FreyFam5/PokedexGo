package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/FreyFam5/go/pokedexcli/internal/pokecache"
)

var mapPrefix string = "https://pokeapi.co/api/v2/location-area/"
var cache = pokecache.NewCache(5 * time.Second)

// Cycles the map display forward
func commandMap(c *config) error {
	if c.Next == "" {
		c.Next = mapPrefix
	}

	body, err := getBodyWithCache(c.Next)
	fmt.Printf("This is the body: %v", body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, c); err != nil {
		return err
	}
	if len(c.Results) == 0 {
		fmt.Println("This is the end of the list!")
		return nil
	}

	for idx := range c.Results {
		fmt.Println(c.Results[idx].Name)
	}

	return nil
}

// Cycles the map display backwards
func commandMapB(c *config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	body, err := getBodyWithCache(*c.Previous)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, c); err != nil {
		return err
	}
	for idx := range c.Results {
		fmt.Println(c.Results[idx].Name)
	}

	return nil
}

// Checks the cache to see if the given url's data already exists and returns it if it was found, otherwise it will return the url's body directly from search
func getBodyWithCache(url string) ([]byte, error) {
	var err error
	entry, exists := cache.Get(url)
	if !exists {
		entry, err = getBody(url)
		if err != nil {
			return []byte{}, err
		}
		cache.Add(url, entry)
	}
	return entry, nil
}

// Gets the urls body from search
func getBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("response failed with status | code: %d | body: %s |", res.StatusCode, body)
	}
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
