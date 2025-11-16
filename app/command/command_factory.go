package command

import (
	"os"
	"os/exec"
	"strings"
)

// CommandBuilder is a function that takes a commandline struct and returns a Command
type CommandBuilder func(params parameters) command

// builtinCommands is a map of command names to their respective CommandBuilder
// It's populated by the init() function of each command
var builtinCommands = make(map[string]CommandBuilder)

func Factory(input string) []Command {
	parametersList := newParser(input)
	var commandList []Command

	for _, params := range parametersList {
		// Switch case to determine which command to return
		if builder, exists := builtinCommands[params.name]; exists {
			commandList = append(commandList, Command{builder(params), params})
		}

		_, err := exec.LookPath(params.name)
		if err == nil {
			// Cast command to ExeCommand
			commandList = append(commandList, Command{execommand(params), params})
		}

		// For any other case, return an UnknownCommand
		// Cast command to UnknwonCommand
		commandList = append(commandList, Command{unknowncommand(params), params})
	}

	return commandList
}

// NewParser creates a new parser and immediately parses the input
func newParser(input string) []parameters {
	var parametersArray []parameters
	tokensArray := createTokens(input)
	for _, tokens := range tokensArray {
		params := createParams(tokens)
		parametersArray = append(parametersArray, params)
	}
	return parametersArray
}

// Return an array of array.
// The first array correspond to a list of commands.
// The second array is the list of arguments for each command
func createTokens(input string) [][]string {
	var tokens []string
	var commandSplit [][]string

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
		case char == ';' && !inSingleQuote && !inDoubleQuote:
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
			}
			commandSplit = append(commandSplit, tokens)
			// Clear the current tokens array as the next tokens belongs to the next command
			tokens = nil
			currentToken.Reset()
		default:
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
		commandSplit = append(commandSplit, tokens)
	}

	return commandSplit
}

func createParams(tokens []string) parameters {
	var params parameters
	params.name = tokens[0]
	params.stdin = os.Stdin
	params.stdout = os.Stdout
	params.stderr = os.Stderr

	// The user input is just one token
	// We just have a command name
	if len(tokens) == 1 {
		return params
	}

	i, token := findRedirectToken(tokens)

	params.args = tokens[1:i]

	// The user doesn't give a filepath after the redirection
	// We keep the stdout as it is
	if len(tokens) == i {
		return params
	}

	filepath := tokens[i+1]

	setStdout(token, params, filepath)

	return params
}

func findRedirectToken(tokens []string) (int, string) {
	for i, token := range tokens {
		if (len(token) == 1 && token[0] == '>') ||
			(len(token) == 2 && token[1] == '>') {
			return i, token
		}
	}
	return len(tokens), ""
}

func setStdout(token string, params parameters, filepath string) {
	switch token {
	case ">", "1>":
		params.stdout, _ = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	case ">>", "1>>":
		params.stdout, _ = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	case "2>":
		params.stderr, _ = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	case "2>>":
		params.stderr, _ = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	}
}
