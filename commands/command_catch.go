package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/rahulkumarpahwa/pokedexcli/cache"
	"github.com/rahulkumarpahwa/pokedexcli/config"
)

func CommandCatch(config *config.Config, cache *cache.Cache, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("args can't be empty")
	}
	arg := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)

	finalUrl := "https://pokeapi.co/api/v2/pokemon/" + arg

	pokemon, err := CatchRequest(finalUrl, cache, args)
	if err != nil {
		return err
	}

	exp := pokemon.BaseExperience

	// fmt.Println(exp)
	num := rand.Intn(exp + 10)
	// fmt.Println(num)

	if num >= exp {

		fmt.Printf("%s was caught!\n", arg)
		config.Caught[arg] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", arg)
	}

	return nil
}

func CatchRequest(url string, cache *cache.Cache, args []string) (config.Pokemon, error) {

	var pokemon config.Pokemon
	data, ok := cache.Get(url)
	if !ok {
		// slog.Info("fetching using the http request")
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return config.Pokemon{}, err
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return config.Pokemon{}, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return config.Pokemon{}, err
		}
		cache.Add(url, body)
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			return config.Pokemon{}, err
		}

	} else {
		// slog.Info("fetching using the cache")
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return config.Pokemon{}, err
		}
	}

	return pokemon, nil
}
