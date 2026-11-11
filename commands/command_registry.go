package commands

import (
	"github.com/rahulkumarpahwa/pokedexcli/cache"
	"github.com/rahulkumarpahwa/pokedexcli/config"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *config.Config, cache *cache.Cache, args []string) error
}

var Registry map[string]CliCommand
