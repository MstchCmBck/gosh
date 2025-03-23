package command

import (
	"fmt"
	"os/exec"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

// No need to implement init for this kind of not builtin command

type ExeCommand commandline

func (e ExeCommand) Execute() error {
	out, err := exec.Command(e.name, e.args...).Output()
	fmt.Println(string(out))
	return err
}
