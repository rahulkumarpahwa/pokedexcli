package commands

import (
	"fmt"

	"github.com/rahulkumarpahwa/pokedexcli/cache"
	"github.com/rahulkumarpahwa/pokedexcli/config"
)

func CommandPokedex(config *config.Config, cache *cache.Cache, args []string) error {
	fmt.Println("Your Pokedex:")
	for key:=  range config.Caught{
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
