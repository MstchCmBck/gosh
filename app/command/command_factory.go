package command

import (
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/parser"
)

// CommandBuilder is a function that takes a string and returns a Command
type CommandBuilder func(params []string) Command

// builtinCommands is a map of command names to their respective CommandBuilder
// It's populated by the init() function of each command
var builtinCommands = make(map[string]CommandBuilder)

func Factory(command *parser.Parser) Command {
	// Switch case to determine which command to return
	if builder, exists := builtinCommands[command.Name]; exists {
		return builder(command.Args)
	}

	_, err := exec.LookPath(command.Name)
	if err == nil {
		return ExeCommand{command.Name, command.Args}
	}

	// For any other case, return an UnknownCommand
	return UnknownCommand(command.Name)
}
