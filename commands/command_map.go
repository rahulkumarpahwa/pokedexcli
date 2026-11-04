package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rahulkumarpahwa/pokedexcli/config"
)

type PokiLocation struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func CommandMap(config *config.Config) error {
	if config.NextUrl == "" {
		return fmt.Errorf("url does not exists")
	}
	return Request(config.NextUrl, config)
}

func CommandMapB(config *config.Config) error {
	if config.PrevUrl == "" {
		return fmt.Errorf("url does not exists")
	}
	return Request(config.PrevUrl, config)
}

func Request(url string, config *config.Config) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var pokiLocation PokiLocation
	json.NewDecoder(resp.Body).Decode(&pokiLocation)
	if pokiLocation.Next != "" {
		config.NextUrl = pokiLocation.Next
	}
	if pokiLocation.Previous != "" {
		config.PrevUrl = pokiLocation.Previous
	}

	if pokiLocation.Results != nil {
		for _, locations := range pokiLocation.Results {
			fmt.Println(locations.Name)
		}
	}

	return nil
}
