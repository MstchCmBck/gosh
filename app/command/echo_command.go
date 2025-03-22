package command

import (
	"fmt"
)

// EchoCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type EchoCommand []string

func init() {
	builtinCommands["echo"] = func(params []string) Command {
		return EchoCommand(params)
	}
}

func (e EchoCommand) Execute() error {
	for _, arg := range e {
		fmt.Print(arg + " ")
	}
	fmt.Println()
	return nil
}
