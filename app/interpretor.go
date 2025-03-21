package main

import (
	"bufio"
	"fmt"
	"os"
)

func run() {
	for {
		// Print the prompt
		fmt.Fprint(os.Stdout, "$ ")
		// Read user input
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		// Print default message
		fmt.Println(unknownCommand(command))
	}
}

func unknownCommand(command string) string {
	return command[:len(command)-1] + ": command not found"
}
