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
	if builder, exists := builtinCommands[command.GetCommand()]; exists {
		return builder(command.GetArgs())
	}

	_, err := exec.LookPath(command.GetCommand())
	if err == nil {
		return ExeCommand{command.GetCommand(), command.GetArgs()}
	}

	// For any other case, return an UnknownCommand
	return UnknownCommand(command.GetCommand())
}
