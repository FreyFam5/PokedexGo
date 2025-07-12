package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var mapPrefix string = "https://pokeapi.co/api/v2/location-area/"

func commandMap(c *Config) error {
	if !strings.HasPrefix(c.Next, mapPrefix) {
		c.Next = mapPrefix
	}

	res, err := http.Get(c.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	c.Previous = c.Next
	c.Offset += 20
	c.Next = mapPrefix + fmt.Sprintf("?offset%d", c.Offset)

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