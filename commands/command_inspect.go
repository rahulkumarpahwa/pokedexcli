package commands

import (
	"fmt"

	"github.com/rahulkumarpahwa/pokedexcli/cache"
	"github.com/rahulkumarpahwa/pokedexcli/config"
)

func CommandInspect(config *config.Config, cache *cache.Cache, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("args can't be empty")
	}
	arg := args[1]

	pokemon, ok := config.Caught[arg]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)

	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ty := range pokemon.Types {
		fmt.Printf(" - %s\n", ty.Type.Name)
	}

	return nil
}
