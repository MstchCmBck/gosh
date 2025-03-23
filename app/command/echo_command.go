package command

import "fmt"

// EchoCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type EchoCommand []string

func init() {
	builtinCommands["echo"] = func(params []string) Command {
		return EchoCommand(params)
	}
}

func (e EchoCommand) Execute() error {
	var output string
	for _, arg := range e {
		output += arg + " "
	}
	fmt.Println(output)
	return nil
}
