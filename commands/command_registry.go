package commands

import "github.com/rahulkumarpahwa/pokedexcli/config"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *config.Config) error
}

var Registry map[string]CliCommand
