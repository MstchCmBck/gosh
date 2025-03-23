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

// This function is used by all the built-in function to print their result
func print(output string, cmd commandline) {
	switch cmd.redirection {
	case noredirection:
		fmt.Print(output)
	case stdout:
		file, err := os.OpenFile(cmd.filepath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
		}

		_, err = file.WriteString(output)
		if err != nil {
			fmt.Println(err)
		}

		if err = file.Close(); err != nil {
			fmt.Println(err)
		}
	}
}
