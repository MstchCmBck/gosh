package command

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// CommandBuilder is a function that takes a string and returns a Command
type CommandBuilder func(params []string) Command

// builtinCommands is a map of command names to their respective CommandBuilder
// It's populated by the init() function of each command
var builtinCommands = make(map[string]CommandBuilder)

func Factory(input string) Command {
	command := newParser(input)

	// Switch case to determine which command to return
	if builder, exists := builtinCommands[command.name]; exists {
		return builder(command.args)
	}

	_, err := exec.LookPath(command.name)
	if err == nil {
		return ExeCommand{command.name, command.args}
	}

	// For any other case, return an UnknownCommand
	return UnknownCommand(command.name)
}

// NewParser creates a new parser and immediately parses the input
func newParser(input string) commandline {
	var command commandline
	tokens := createTokens(input)
	command.name = tokens[0] // Simple example, adjust as needed
	command.redirection = noredirection
	command.args = []string{}
	if len(tokens) < 2 {
		return command
	}
	var index int
	var err error
	command.redirection, index, err = getRedirectionToken(tokens)
	if err != nil {
		fmt.Println(err)
	}

	if command.redirection != noredirection {
		command.filepath = tokens[index+1]
		command.args = tokens[1:index]
	} else {
		command.args = tokens[1:]
	}

	return command
}

func createTokens(input string) []string {
	var tokens []string

	var currentToken strings.Builder
	inSingleQuote := false
	inDoubleQuote := false
	escapeNext := false

	for i, char := range input {
		switch {
		case escapeNext:
			currentToken.WriteRune(char)
			escapeNext = false
		case char == '\\' && !inDoubleQuote && !inSingleQuote:
			escapeNext = true
		case char == '\\' && inDoubleQuote:
			if input[i+1] == '"' || input[i+1] == '\\' || input[i+1] == '$' {
				escapeNext = true
			} else {
				currentToken.WriteRune(char)
			}
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

func getRedirectionToken(tokens []string) (redirection, int, error) {
	// Implement redirection parsing
	redirection := noredirection
	redirectionIndex := len(tokens)

	for i, token := range tokens {
		if token == ">" || token == "1>" {
			redirection = stdout
			redirectionIndex = i
			break
		} else if token == ">>" {
			redirection = stdoutappend
			redirectionIndex = i
			break
		} else if token == "2>" {
			redirection = stderr
			redirectionIndex = i
			break
		}
	}

	if redirectionIndex == len(tokens)-1 && redirection != noredirection {
		return noredirection, redirectionIndex, errors.New("no file specified for redirection")
	}

	return redirection, redirectionIndex, nil
}
