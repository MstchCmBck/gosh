package command

import (
	"fmt"
	"os"
	"os/exec"
)

// Each new Command must implement the following functions:
// - init() function to register the command in the builtinCommands map
// - Execute() method to execute the command

// No need to implement init for this kind of not builtin command

type execommand commandline

func (e execommand) Execute() error {
	out := exec.Command(e.name, e.args...)
	out.Stdin = os.Stdin
	outstream := e.selectStdoutStream()
	out.Stdout = outstream
	errstream := e.selectStderrStream()
	out.Stderr = errstream
	err := out.Run()
	e.closeStream(outstream)
	e.closeStream(errstream)
	return err
}

func (e execommand) selectStdoutStream() *os.File {
	switch e.redirection {
	case stdout:
		file, _ := os.OpenFile(e.filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		return file
	case stdoutappend:
		file, _ := os.OpenFile(e.filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		return file
	default:
		return os.Stdout
	}
}

func (e execommand) selectStderrStream() *os.File {
	switch e.redirection {
	case stderr:
		file, _ := os.OpenFile(e.filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		return file
	case stderrappend:
		file, _ := os.OpenFile(e.filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		return file
	default:
		return os.Stdout
	}
}

func (e execommand) closeStream(stream *os.File) {
	// Only close if it's not Stdout or Stderr
	if stream != os.Stdout && stream != os.Stderr {
		if err := stream.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
