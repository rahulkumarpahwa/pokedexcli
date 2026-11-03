package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		words := cleanInput(text)
		fmt.Printf("Your command was: %s\n", words[0])
		fmt.Print("Pokedex > ")
	}
}
