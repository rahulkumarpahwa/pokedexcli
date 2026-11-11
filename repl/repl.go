package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rahulkumarpahwa/pokedexcli/cache"
	"github.com/rahulkumarpahwa/pokedexcli/commands"
	"github.com/rahulkumarpahwa/pokedexcli/config"
)

func cleanInput(text string) []string {
	small := strings.ToLower(text)
	return strings.Fields(small)
}

func StartRepl(cfg *config.Config) {
	cache := cache.NewCache(5 * time.Minute)

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Err() != nil {
		fmt.Printf("scanner error: %v\n", scanner.Err())
		return
	}

	commands.Registry = map[string]commands.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commands.CommandExit,
		},

		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commands.CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    commands.CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "displays the names of previous 20 location areas in the Pokemon world",
			Callback:    commands.CommandMapB,
		},
		"explore": {
			Name:        "explore <location_name>",
			Description: "list of all the Pokémon located in the location",
			Callback:    commands.CommandExplore,
		},
		"catch": {
			Name:        "catch <pokemon_name>",
			Description: "catching Pokemon adds them to the user's Pokedex",
			Callback:    commands.CommandCatch,
		},
		"inspect" :{
			Name : "inspect <pokemon_name>",
			Description: "will only allow players to see details about a Pokemon if they have seen it before",
			Callback: commands.CommandInspect,
		}, 
	}

	for scanner.Scan() {
		text := scanner.Text()
		words := cleanInput(text)
		for _, word := range words {
			val, ok := commands.Registry[word]
			if !ok {
				// fmt.Println("Unknown command")
				continue
			}
			err := val.Callback(cfg, cache, words)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
