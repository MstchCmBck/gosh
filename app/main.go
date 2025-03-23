package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/command"
)

func main() {
	for {
		// Print the prompt
		fmt.Fprint(os.Stdout, "$ ")
		// Read user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := command.Factory(input)
		cmd.Execute()
	}
}
