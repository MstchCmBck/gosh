package command

import (
	"errors"
	"os"
	"strconv"
)

// exitcommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type exitcommand commandline

func init() {
	builtinCommands["exit"] = func(params commandline) Command {
		return exitcommand(params)
	}
}

func (e exitcommand) Execute() error {
	// If the command has no parameters, exit with code 0
	if len(e.args) == 0 {
		os.Exit(0)
	}
	// Extract the errCode from the command
	errCode, err := strconv.Atoi(e.args[0])
	// If the conversion fails, consider the command as not an exit command
	if err != nil {
		return errors.New("exit command parameter is not a number")
	}
	os.Exit(errCode)
	return nil
}
