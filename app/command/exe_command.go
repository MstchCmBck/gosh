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
		file, _ := os.OpenFile(e.filepath, os.O_RDWR|os.O_CREATE, 0644)
		return file
	case noredirection:
	default:
		return os.Stdout
	}
	return os.Stdout
}

func (e execommand) selectStderrStream() *os.File {
	switch e.redirection {
	case stderr:
		file, _ := os.OpenFile(e.filepath, os.O_RDWR|os.O_CREATE, 0644)
		return file
	case noredirection:
	default:
		return os.Stdout
	}
	return os.Stdout
}

func (e execommand) closeStream(stream *os.File) {
	switch e.redirection {
	case stdout:
		if err := stream.Close(); err != nil {
			fmt.Println(err)
		}
		return
	case noredirection:
	default:
		return
	}
}
