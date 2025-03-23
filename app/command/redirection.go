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
	case stdoutappend:
		appendFile(input, cmd.filepath)
	default:
		fmt.Print(input)
	}
}

// This function is used by all the built-in function to printOut their result
func printErr(input string, cmd commandline) {
	switch cmd.redirection {
	case stderr:
		printFile(input, cmd.filepath)
	case stderrappend:
		appendFile(input, cmd.filepath)
	default:
		fmt.Print(input)
	}
}

func printFile(input string, filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

func appendFile(input string, filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
