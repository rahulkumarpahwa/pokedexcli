package main

import (
	"github.com/rahulkumarpahwa/pokedexcli/config"
	"github.com/rahulkumarpahwa/pokedexcli/repl"
)

func main() {
	cfg := config.Config{
		ExploreUrl: "https://pokeapi.co/api/v2/location-area/",
		NextUrl: "https://pokeapi.co/api/v2/location-area/",
		Caught:  make(map[string]config.Pokemon),
	}
	repl.StartRepl(&cfg)
}
