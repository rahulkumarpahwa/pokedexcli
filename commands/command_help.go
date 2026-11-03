package commands

import (
	"fmt"
)

func CommandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for key, value := range Registry {
		fmt.Printf("%s: %s\n", key, value.Description)
	}

	return nil
}
