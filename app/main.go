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
		_, err := cmd.Execute()
		if err != nil {
			fmt.Print(err)
		}

		// TODO: This should be handle by the command itself
		// To remove soon
		// if prsr.Redirection == parser.Stdout {
		// 	os.WriteFile(prsr.Filepath, []byte(output), 0644)
		// } else {
		// 	fmt.Println(output)
		// }

	}
}
