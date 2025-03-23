package command

import (
	"fmt"
	"os/exec"
)

// TypeCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type TypeCommand []string

func init() {
	builtinCommands["type"] = func(params []string) Command {
		return TypeCommand(params)
	}
}

func (t TypeCommand) Execute() error {
	var output string
	_, find := builtinCommands[string(t[0])]
	if find {
		output = string(t[0]) + " is a shell builtin"
		fmt.Println(output)
		return nil
	}

	path, err := exec.LookPath(string(t[0]))
	if err != nil {
		output = string(t[0]) + ": not found"
		fmt.Println(output)
		return nil
	}

	output = string(t[0]) + " is " + path
	fmt.Println(output)
	return nil
}
