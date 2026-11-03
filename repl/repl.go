package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rahulkumarpahwa/pokedexcli/commands"
)

func cleanInput(text string) []string {
	small := strings.ToLower(text)
	return strings.Fields(small)
}

func StartRepl() {
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
			err := val.Callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
