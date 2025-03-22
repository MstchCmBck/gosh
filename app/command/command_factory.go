package command

import (
	"os/exec"

	"github.com/mattn/go-shellwords"
)

// CommandBuilder is a function that takes a string and returns a Command
type CommandBuilder func(params []string) Command

// builtinCommands is a map of command names to their respective CommandBuilder
// It's populated by the init() function of each command
var builtinCommands = make(map[string]CommandBuilder)

func Factory(command string) Command {
	// Extract the command name and its parameters
	splittedcommand, err := shellwords.Parse(command)
	if err != nil {
		return UnknownCommand(command)
	}
	commandName := splittedcommand[0]
	parameters := splittedcommand[1:]

	// Switch case to determine which command to return
	if builder, exists := builtinCommands[commandName]; exists {
		return builder(parameters)
	}

	_, err = exec.LookPath(commandName)
	if err == nil {
		return ExeCommand{commandName, parameters}
	}

	// For any other case, return an UnknownCommand
	return UnknownCommand(commandName)
}
