package command

import (
	"fmt"
	"os"

	"github.com/chzyer/readline"
)

// pwdcommand implements the "pwd" command, which prints the current working directory.
// It retrieves the current working directory using os.Getwd() and prints it to standard output.
// It also registers itself in the builtinCommands map and adds "pwd" to the completion suggestions.
// pwdcommand implements the Command interface.
type pwdcommand parameters

// init registers the pwdcommand in the builtinCommands map and adds "pwd" to the completion suggestions.
func init() {
	builtinCommands["pwd"] = func(params parameters) command {
		return pwdcommand(params)
	}

	BuiltinCompletion = append(BuiltinCompletion, readline.PcItem("pwd"))
}

// execute retrieves the current working directory and prints it to standard output.
func (p pwdcommand) execute() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	wd += "\n"
	fmt.Fprint(p.stdout, wd)

	return nil
}
