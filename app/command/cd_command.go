package command

import (
	"fmt"
	"os"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

type cdcommand parameters

func init() {
	builtinCommands["cd"] = func(params parameters) command {
		return cdcommand(params)
	}
}

func (c cdcommand) execute() error {
	var err error
	dir := string(c.args[0])
	if c.args[0] == "~" {
		dir, err = os.UserHomeDir()
		if err != nil {
			return err
		}
	}
	err = os.Chdir(dir)
	var message string
	if err != nil {
		message = "cd: " + dir + ": No such file or directory\n"
	}
	fmt.Fprint(c.stdout, message)

	return err
}
