package command

import "os"

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

func init() {
	builtinCommands["pwd"] = func(params []string) Command {
		return PwdCommand(params)
	}
}

type PwdCommand []string

func (p PwdCommand) Execute() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	println(wd)
	return nil
}
