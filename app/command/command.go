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
	cmd    command
	params parameters
}

func (c Command) Execute() error {
	defer close(c.params)
	return c.cmd.execute()
}

type parameters struct {
	name   string
	args   []string
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

type command interface {
	execute() error
}

func close(params parameters) error {
	var err_stdin error
	var err_stdout error
	var err_stderr error

	if closer, ok := params.stdin.(io.Closer); ok && params.stdin != os.Stdin {
		err_stdin = closer.Close()
	}
	if closer, ok := params.stdout.(io.Closer); ok && params.stdout != os.Stdout {
		err_stdout = closer.Close()
	}
	if closer, ok := params.stderr.(io.Closer); ok && params.stderr != os.Stderr {
		err_stderr = closer.Close()
	}

	if err_stdin != nil {
		return err_stdin
	}
	if err_stdout != nil {
		return err_stdout
	}
	return err_stderr
}
