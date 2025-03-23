package command

import "fmt"

// UnknownCommand string contains the command name that is unknown
// It implement the Execute method of the Command interface
type UnknownCommand commandline

// No need to register the UnknownCommand in the builtinCommands map as it's the default command

func (u UnknownCommand) Execute() error {
	unknown := string(u.name)
	fmt.Println(unknown + ": command not found")
	return nil
}
