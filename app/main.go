package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/command"
	"github.com/codecrafters-io/shell-starter-go/app/parser"
)

func main() {
	for {
		// Print the prompt
		fmt.Fprint(os.Stdout, "$ ")
		// Read user input
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		prsr := parser.NewParser(userInput)
		cmd := command.Factory(prsr)
		output, err := cmd.Execute()
		if err != nil {
			fmt.Print(err)
		}

		if prsr.GetRedirection() == parser.Stdout {
			os.WriteFile(prsr.GetFilepath(), []byte(output), 0644)
		} else {
			fmt.Println(output)
		}

	}
}
