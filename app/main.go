package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		// Print the prompt
		fmt.Fprint(os.Stdout, "$ ")
		// Read user input
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command := commandFactory(userInput)
		if err := command.Execute(); err != nil {
			fmt.Println(err)
		}
	}
}
