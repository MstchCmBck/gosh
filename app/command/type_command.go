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
	_, find := builtinCommands[string(t[0])]
	if find {
		fmt.Println(t[0] + " is a shell builtin")
		return nil
	}

	path, err := exec.LookPath(string(t[0]))
	if err != nil {
		fmt.Println(t[0] + ": not found")
		return nil
	}

	fmt.Println(string(t[0]) + " is " + path)
	return nil
}
