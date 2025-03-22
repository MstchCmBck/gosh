package command

import (
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

func (t TypeCommand) Execute() (string, error) {
	var output string
	_, find := builtinCommands[string(t[0])]
	if find {
		output = string(t[0]) + " is a shell builtin"
		return output, nil
	}

	path, err := exec.LookPath(string(t[0]))
	if err != nil {
		output = string(t[0]) + ": not found"
		return output, nil
	}

	output = string(t[0]) + " is " + path
	return output, nil
}
