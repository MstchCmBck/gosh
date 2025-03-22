package command

import (
	"fmt"
	"os/exec"
	"strings"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

// No need to implement init for this kind of not builtin command

type ExeCommand struct {
	commandName string
	parameters  string
}

func (e ExeCommand) Execute() error {
	fields := strings.Fields(e.parameters)
	out, err := exec.Command(e.commandName, fields...).Output()
	fmt.Print(string(out))
	return err
}
