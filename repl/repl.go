package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rahulkumarpahwa/pokedexcli/commands"
	"github.com/rahulkumarpahwa/pokedexcli/config"
)

func cleanInput(text string) []string {
	small := strings.ToLower(text)
	return strings.Fields(small)
}

func StartRepl(cfg *config.Config) {
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
		"mapb" : {
			Name: "mapb",
			Description: "displays the names of previous 20 location areas in the Pokemon world",
			Callback: commands.CommandMapB,
		},
	}

	for scanner.Scan() {
		text := scanner.Text()
		words := cleanInput(text)
		for _, word := range words {
			val, ok := commands.Registry[word]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			err := val.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
