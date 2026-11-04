package commands

import (
	"fmt"
	"os"

	"github.com/rahulkumarpahwa/pokedexcli/config"
)

func CommandExit(config *config.Config) error {

	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)

	return nil
}
