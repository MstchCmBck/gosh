package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func run() int {
	for {
		// Print the prompt
		fmt.Fprint(os.Stdout, "$ ")
		// Read user input
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if exit, errCode := isExit(command); exit {
			// Exit the program with the given errCode
			return errCode
		}
		// Print default message
		fmt.Println(unknownCommand(command))
	}
}

func isExit(command string) (bool, int) {
	// Will match "exit 123", not match "exit" or "exit123"
	re := regexp.MustCompile(`^exit\s(\d+)`)	
	// Early return if the command does not match the regex
	if !re.Match([]byte(command)) {
		return false, 0
	}
	// Extract the errCode from the command
	errCode, err := strconv.Atoi(string(re.FindSubmatch([]byte(command))[1]))
	// If the conversion fails, consider the command as not an exit command
	if err != nil {
		return false, 0
	}
	return true, errCode
}

func unknownCommand(command string) string {
	return command[:len(command)-1] + ": command not found"
}
