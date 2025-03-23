package command

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

type commandline struct {
	name        string
	args        []string
	redirection redirection
	filepath    string
}

type Command interface {
	Execute() error
}
