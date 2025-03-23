package command

import (
	"os"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

type cdcommand commandline

func init() {
	builtinCommands["cd"] = func(params commandline) Command {
		return cdcommand(params)
	}
}

func (c cdcommand) Execute() error {
	var err error
	dir := string(c.args[0])
	if c.args[0] == "~" {
		dir, err = os.UserHomeDir()
		if err != nil {
			return err
		}
	}
	err = os.Chdir(dir)
	if err != nil {
		print("cd: "+dir+": No such file or directory\n", commandline(c))
	}
	return err
}
