package repl

import "strings"

func cleanInput(text string) []string {
	small := strings.ToLower(text)
	return strings.Fields(small)
}
