package command

import (
	"errors"
)

// UnknownCommand string contains the command name that is unknown
// It implement the Execute method of the Command interface
type UnknownCommand string

// No need to register the UnknownCommand in the builtinCommands map as it's the default command

func (u UnknownCommand) Execute() (string, error) {
	return "", errors.New(string(u) + ": command not found")
}
