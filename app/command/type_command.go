package command

import (
	"fmt"
	"os/exec"
)

// TypeCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type TypeCommand string

func init() {
	builtinCommands["type"] = func(params string) Command {
		return TypeCommand(params)
	}
}

func (t TypeCommand) Execute() error {
	_, find := builtinCommands[string(t)]
	if find {
		fmt.Println(t + " is a shell builtin")
		return nil
	}

	path, err := exec.LookPath(string(t))
	if err != nil {
		fmt.Println(t + ": not found")
		return nil
	}

	fmt.Println(string(t) + " is " + path + "/" + string(t))
	return nil
}
