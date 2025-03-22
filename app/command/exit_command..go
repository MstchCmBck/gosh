package command

import (
	"errors"
	"os"
	"regexp"
	"strconv"
)

// ExitCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type ExitCommand string

func init() {
	builtinCommands["exit"] = func(params string) Command {
		return ExitCommand(params)
	}
}

func (e ExitCommand) Execute() error {
	// If the command has no parameters, exit with code 0
	if e == "" {
		os.Exit(0)
	}
	// Will match "exit 123", not match "exit" or "exit123"
	re := regexp.MustCompile(`^(\d+).*`)
	// Extract the errCode from the command
	errCode, err := strconv.Atoi(string(re.FindSubmatch([]byte(e))[1]))
	// If the conversion fails, consider the command as not an exit command
	if err != nil {
		return errors.New("exit command parameter is not a number")
	}
	os.Exit(errCode)
	return nil
}
