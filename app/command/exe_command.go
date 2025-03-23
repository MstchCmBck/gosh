package command

import (
	"os/exec"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

// No need to implement init for this kind of not builtin command

type execommand commandline

func (e execommand) Execute() error {
	out := exec.Command(e.name, e.args...)

	out.Stdin = e.stdin
	out.Stderr = e.stderr
	out.Stdout = e.stdout

	err := out.Run()
	return err
}
