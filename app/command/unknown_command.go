package command

// unknowncommand string contains the command name that is unknown
// It implement the Execute method of the Command interface
type unknowncommand commandline

// No need to register the UnknownCommand in the builtinCommands map as it's the default command

func (u unknowncommand) Execute() error {
	unknown := string(u.name)
	output := unknown + ": command not found"
	print(output, commandline(u))
	return nil
}
