package command

import (
	"io"
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
	parametersList := createParametersPerCommand(input)
	var commandList []Command

	for _, params := range parametersList {
		var cmd command
		var err error

		if builder, exists := builtinCommands[params.name]; exists {
			cmd = builder(params)
		} else if _, err = exec.LookPath(params.name); err == nil {
			cmd = execommand(params)
		} else {
			cmd = unknowncommand(params)
		}

		commandList = append(commandList, Command{cmd, params})
	}

	return commandList
}

func createParametersPerCommand(input string) []parameters {
	var tokens []string
	var parametersList []parameters

	var currentToken strings.Builder
	inSingleQuote := false
	inDoubleQuote := false
	escapeNext := false
	var nextStdin io.Reader = os.Stdin

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
			params := createParams(tokens, nextStdin)
			nextStdin = os.Stdin
			parametersList = append(parametersList, params)
			// Clear the current tokens array as the next tokens belongs to the next command
			tokens = nil
			currentToken.Reset()
		case char == '|' && !inSingleQuote && !inDoubleQuote:
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
			}
			params := createParams(tokens, nextStdin)

			r, w := io.Pipe()
			params.stdout = w
			nextStdin = r

			parametersList = append(parametersList, params)
			tokens = nil
			currentToken.Reset()
		default:
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
		params := createParams(tokens, nextStdin)
		parametersList = append(parametersList, params)
	}

	return parametersList
}

func createParams(tokens []string, stdin io.Reader) parameters {
	var params parameters
	params.name = tokens[0]
	params.stdin = stdin
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
