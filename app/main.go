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
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := command.Factory(userInput)
		var output string
		var err error
		if output, err = cmd.Execute(); err != nil {
			fmt.Println(err)
		}
		fmt.Print(output)
	}
}
