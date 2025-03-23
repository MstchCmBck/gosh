package command

import (
	"fmt"
	"os"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

type PwdCommand commandline

func init() {
	builtinCommands["pwd"] = func(params commandline) Command {
		return PwdCommand(params)
	}
}

func (p PwdCommand) Execute() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(wd)
	return nil
}
