package command

import (
	"errors"
	"os"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

func init() {
	builtinCommands["cd"] = func(params string) Command {
		return CdCommand(params)
	}
}

type CdCommand string

func (c CdCommand) Execute() error {
	err := os.Chdir(string(c))
	if err != nil {
		err = errors.New("cd: " + string(c) + ": No such file or directory")
	}
	return err
}
