package command

import (
	"fmt"
	"os"
)

type redirection int

const (
	noredirection redirection = iota
	stdout
	stdoutappend
	stderr
	stderrappend
)

// This function is used by all the built-in function to printOut their result
func printOut(input string, cmd commandline) {
	switch cmd.redirection {
	case stdout:
		printFile(input, cmd.filepath)
	default:
		fmt.Print(input)
	}
}

// This function is used by all the built-in function to printOut their result
func printErr(input string, cmd commandline) {
	switch cmd.redirection {
	case stderr:
		printFile(input, cmd.filepath)
	default:
		fmt.Print(input)
	}
}

func printFile(input string, filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.WriteString(input)
	if err != nil {
		fmt.Println(err)
	}

	if err = file.Close(); err != nil {
		fmt.Println(err)
	}
}
