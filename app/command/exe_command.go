package command

import (
	"os/exec"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

// No need to implement init for this kind of not builtin command

type ExeCommand struct {
	commandName string
	parameters  []string
}

func (e ExeCommand) Execute() (string, error) {
	out, err := exec.Command(e.commandName, e.parameters...).Output()
	return string(out), err
}
