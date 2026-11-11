package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rahulkumarpahwa/pokedexcli/cache"
	"github.com/rahulkumarpahwa/pokedexcli/config"
)

type ExplorePokimon struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	GameIndex         int    `json:"game_index"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func CommandExplore(config *config.Config, cache *cache.Cache, args []string) error {
	if config.ExploreUrl == "" {
		return fmt.Errorf("next url does not exists")
	}
	return ExploreRequest(config.ExploreUrl, cache, args)
}

func ExploreRequest(url string, cache *cache.Cache, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("args can't be empty")
	}

	finalUrl := url + args[1]

	var explorePokimon ExplorePokimon
	data, ok := cache.Get(finalUrl)
	if !ok {
		// slog.Info("fetching using the http request")
		req, err := http.NewRequest("GET", finalUrl, nil)
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
		err = json.Unmarshal(body, &explorePokimon)
		if err != nil {
			return err
		}

	} else {
		// slog.Info("fetching using the cache")
		err := json.Unmarshal(data, &explorePokimon)
		if err != nil {
			return err
		}
	}

	if explorePokimon.PokemonEncounters != nil {
		for _, pokimon := range explorePokimon.PokemonEncounters {
			fmt.Println(pokimon.Pokemon.Name)
		}
	}

	return nil
}
