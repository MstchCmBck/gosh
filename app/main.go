package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Read user input
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// Print default message
	fmt.Println(unknownCommand(command))
}
