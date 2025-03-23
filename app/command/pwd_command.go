package command

import (
	"os"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

type pwdcommand commandline

func init() {
	builtinCommands["pwd"] = func(params commandline) Command {
		return pwdcommand(params)
	}
}

func (p pwdcommand) Execute() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	wd += "\n"
	print(wd, commandline(p))
	return nil
}
