package command

import (
	"os"
	"os/exec"
	"strings"
)

// CommandBuilder is a function that takes a string and returns a Command
type CommandBuilder func(params commandline) Command

// builtinCommands is a map of command names to their respective CommandBuilder
// It's populated by the init() function of each command
var builtinCommands = make(map[string]CommandBuilder)

func Factory(input string) Command {
	command := newParser(input)

	// Switch case to determine which command to return
	if builder, exists := builtinCommands[command.name]; exists {
		return builder(command)
	}

	_, err := exec.LookPath(command.name)
	if err == nil {
		// Cast command to ExeCommand
		return execommand(command)
	}

	// For any other case, return an UnknownCommand
	// Cast command to UnknwonCommand
	return unknowncommand(command)
}

// NewParser creates a new parser and immediately parses the input
func newParser(input string) commandline {
	tokens := createTokens(input)
	cmd := createCommand(tokens)
	return cmd
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

func createCommand(tokens []string) commandline {
	var cmd commandline
	cmd.name = tokens[0]
	cmd.stdin = os.Stdin
	cmd.stdout = os.Stdout
	cmd.stderr = os.Stderr

	// The user input is just one token
	// We just have a command name
	if len(tokens) == 1 {
		return cmd
	}

	i, token := findRedirectToken(tokens)

	cmd.args = tokens[1:i]

	// The user doesn't give a filepath after the redirection
	// We keep the stdout as it is
	if len(tokens) == i {
		return cmd
	}

	filepath := tokens[i+1]

	setStdout(token, cmd, filepath)

	return cmd
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

func setStdout(token string, cmd commandline, filepath string) {
	switch token {
	case ">", "1>":
		cmd.stdout, _ = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	case ">>", "1>>":
		cmd.stdout, _ = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	case "2>":
		cmd.stderr, _ = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	case "2>>":
		cmd.stderr, _ = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	}
}
