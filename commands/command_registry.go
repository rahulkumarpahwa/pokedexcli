package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var Registry map[string]CliCommand
