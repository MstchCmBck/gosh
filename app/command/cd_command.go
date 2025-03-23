package command

import (
	"errors"
	"os"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

func init() {
	builtinCommands["cd"] = func(params []string) Command {
		return CdCommand(params)
	}
}

type CdCommand []string

func (c CdCommand) Execute() error {
	var err error
	dir := string(c[0])
	if c[0] == "~" {
		dir, err = os.UserHomeDir()
		if err != nil {
			return err
		}
	}
	err = os.Chdir(dir)
	if err != nil {
		err = errors.New("cd: " + string(c[0]) + ": No such file or directory")
	}
	return err
}
