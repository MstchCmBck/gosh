package command

import (
	"errors"
	"os"
	"strconv"
)

// ExitCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type ExitCommand []string

func init() {
	builtinCommands["exit"] = func(params []string) Command {
		return ExitCommand(params)
	}
}

func (e ExitCommand) Execute() (string, error) {
	// If the command has no parameters, exit with code 0
	if len(e) == 0 {
		os.Exit(0)
	}
	// Extract the errCode from the command
	errCode, err := strconv.Atoi(e[0])
	// If the conversion fails, consider the command as not an exit command
	if err != nil {
		return "", errors.New("exit command parameter is not a number")
	}
	os.Exit(errCode)
	return "", nil
}
