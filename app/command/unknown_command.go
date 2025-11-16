package command

import "fmt"

// unknowncommand string contains the command name that is unknown
// It implement the Execute method of the Command interface
type unknowncommand parameters

// No need to register the UnknownCommand in the builtinCommands map as it's the default command

func (u unknowncommand) execute() error {
	unknown := string(u.name)
	output := unknown + ": command not found\n"
	fmt.Fprint(u.stdout, output)

	return nil
}
