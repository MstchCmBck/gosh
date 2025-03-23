package command

import "fmt"

// echocommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type echocommand commandline

func init() {
	builtinCommands["echo"] = func(params commandline) Command {
		return echocommand(params)
	}
}

func (e echocommand) Execute() error {
	var output string
	for _, arg := range e.args {
		output += arg + " "
	}
	output += "\n"
	fmt.Fprint(e.stdout, output)

	close(commandline(e))
	return nil
}
