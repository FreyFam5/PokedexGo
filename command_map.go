package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var mapPrefix string = "https://pokeapi.co/api/v2/location-area/"

// Cycles the map display forward
func commandMap(c *config) error {
	if c.Next == "" {
		c.Next = mapPrefix
	}

	res, err := http.Get(c.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()


	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status | code: %d | body: %s |", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, c); err != nil {
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
	if c.Previous == nil{
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(*c.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status | code: %d | body: %s |", res.StatusCode, body)
	}
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