package command

import (
	"fmt"
	"os/exec"
)

// TypeCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface

type TypeCommand commandline

func init() {
	builtinCommands["type"] = func(params commandline) Command {
		return TypeCommand(params)
	}
}

func (t TypeCommand) Execute() error {
	var output string
	program := string(t.args[0])
	_, find := builtinCommands[program]
	if find {
		output = program + " is a shell builtin"
		fmt.Println(output)
		return nil
	}

	path, err := exec.LookPath(program)
	if err != nil {
		output = program + ": not found"
		fmt.Println(output)
		return nil
	}

	output = program + " is " + path
	fmt.Println(output)
	return nil
}
