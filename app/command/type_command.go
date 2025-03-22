package command

import (
	"fmt"
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
	if !find {
		err := UnknownCommand(t).Execute()
		return err
	}
	fmt.Println(t + " is a shell builtin")
	return nil
}
