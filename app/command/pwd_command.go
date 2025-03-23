package command

import (
	"fmt"
	"os"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

type pwdcommand commandline

func init() {
	builtinCommands["pwd"] = func(params commandline) command {
		return pwdcommand(params)
	}
}

func (p pwdcommand) execute() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	wd += "\n"
	fmt.Fprint(p.stdout, wd)

	return nil
}
