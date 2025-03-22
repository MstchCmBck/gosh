package command

import (
	"os/exec"
	"strings"
)

// CommandBuilder is a function that takes a string and returns a Command
type CommandBuilder func(params []string) Command

// builtinCommands is a map of command names to their respective CommandBuilder
// It's populated by the init() function of each command
var builtinCommands = make(map[string]CommandBuilder)

func Factory(command string) Command {
	// Extract the command name and its parameters
	splittedcommand := parseInput(command)
	commandName := splittedcommand[0]
	parameters := splittedcommand[1:]

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

func parseInput(input string) []string {
	var tokens []string
	var currentToken strings.Builder
	inSingleQuote := false
	inDoubleQuote := false
	escapeNext := false

	for _, char := range input {
		switch {
		case escapeNext:
			currentToken.WriteRune(char)
			escapeNext = false
		case char == '\\' && !inDoubleQuote:
			escapeNext = true
		case char == '\'' && !inDoubleQuote:
			inSingleQuote = !inSingleQuote
		case char == '"' && !inSingleQuote:
			inDoubleQuote = !inDoubleQuote
		case char == ' ' && !inSingleQuote && !inDoubleQuote:
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
		case char == '\n' && !inSingleQuote && !inDoubleQuote:
			continue
		default:
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
