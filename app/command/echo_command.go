package command

import "fmt"

// EchoCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type EchoCommand commandline

func init() {
	builtinCommands["echo"] = func(params commandline) Command {
		return EchoCommand(params)
	}
}

func (e EchoCommand) Execute() error {
	var output string
	for _, arg := range e.args {
		output += arg + " "
	}
	fmt.Println(output)
	return nil
}
