package command

import (
	"fmt"
	"os"

	"github.com/chzyer/readline"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

type pwdcommand parameters

func init() {
	builtinCommands["pwd"] = func(params parameters) command {
		return pwdcommand(params)
	}

	BuiltinCompletion = append(BuiltinCompletion, readline.PcItem("pwd"))
}

func (p pwdcommand) execute() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	wd += "\n"
	fmt.Fprint(p.stdout, wd)

	return nil
}
