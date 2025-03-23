package command

import (
	"fmt"
	"os/exec"
)

// TypeCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface

type typecommand commandline

func init() {
	builtinCommands["type"] = func(params commandline) Command {
		return typecommand(params)
	}
}

func (t typecommand) Execute() error {
	var output string
	program := string(t.args[0])
	_, find := builtinCommands[program]
	if find {
		output = program + " is a shell builtin\n"
		print(output, commandline(t))
		return nil
	}

	path, err := exec.LookPath(program)
	if err != nil {
		output = program + ": not found\n"
		print(output, commandline(t))
		return nil
	}

	output = program + " is " + path
	fmt.Println(output)
	return nil
}
