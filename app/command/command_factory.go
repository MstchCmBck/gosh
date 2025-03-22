package command

import (
	"os/exec"
	"regexp"
	"strings"
)

// CommandBuilder is a function that takes a string and returns a Command
type CommandBuilder func(params string) Command

// builtinCommands is a map of command names to their respective CommandBuilder
// It's populated by the init() function of each command
var builtinCommands = make(map[string]CommandBuilder)

func Factory(command string) Command {
	// Extract the command name and its parameters
	re := regexp.MustCompile(`^(\w+)(.*)`)
	splittedCommand := re.FindSubmatch([]byte(command))
	commandName := string(splittedCommand[1])
	parameters, _ := strings.CutPrefix(string(splittedCommand[2]), " ")

	// Switch case to determine which command to return
	if builder, exists := builtinCommands[commandName]; exists {
		return builder(parameters)
	}

	_, err := exec.LookPath(commandName)
	if err == nil {
		return ExeCommand{commandName, parameters}
	}

	// For any other case, return an UnknownCommand
	return UnknownCommand(commandName)
}
