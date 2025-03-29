package command

import (
	"io"
	"os"

	"github.com/chzyer/readline"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

var BuiltinCompletion = make([]readline.PrefixCompleterInterface, 0)

type Command struct {
	cmd     command
	cmdline commandline
}

func (c Command) Execute() error {
	defer close(c.cmdline)
	return c.cmd.execute()
}

type commandline struct {
	name   string
	args   []string
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

type command interface {
	execute() error
}

func close(cmd commandline) error {
	var err_stdout error
	var err_stderr error
	if closer, ok := cmd.stdout.(io.Closer); ok && cmd.stdout != os.Stdout {
		err_stdout = closer.Close()
	}
	if closer, ok := cmd.stderr.(io.Closer); ok && cmd.stderr != os.Stderr {
		err_stderr = closer.Close()
	}
	if err_stdout != nil {
		return err_stdout
	}
	return err_stderr
}
