package commands

import (
	"encoding/json"
	"fmt"
	"io"
	// "log/slog"
	"net/http"

	"github.com/rahulkumarpahwa/pokedexcli/cache"
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

func CommandMap(config *config.Config, cache *cache.Cache, args []string) error {
	if config.NextUrl == "" {
		return fmt.Errorf("next url does not exists")
	}

	return Request(config.NextUrl, config, cache)
}

func CommandMapB(config *config.Config, cache *cache.Cache, args []string) error {
	if config.PrevUrl == "" {
		return fmt.Errorf("prev url does not exists")
	}
	return Request(config.PrevUrl, config, cache)
}

func Request(url string, config *config.Config, cache *cache.Cache) error {
	var pokiLocation PokiLocation
	data, ok := cache.Get(url)
	if !ok {
		// slog.Info("fetching using the http request")
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
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cache.Add(url, body)
		err = json.Unmarshal(body, &pokiLocation)
		if err != nil {
			return err
		}

	} else {
		// slog.Info("fetching using the cache")
		err := json.Unmarshal(data, &pokiLocation)
		if err != nil {
			return err
		}
	}

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
