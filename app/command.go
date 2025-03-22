package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func commandFactory(command string) Command {
	// Extract the command name and its parameters
	re := regexp.MustCompile(`^(\w+)(.*)`)
	splittedCommand := re.FindSubmatch([]byte(command))
	commandName := string(splittedCommand[1])
	parameters, _ := strings.CutPrefix(string(splittedCommand[2]), " ")

	// Switch case to determine which command to return
	if commandName == "exit" {
		return ExitCommand(parameters)
	}

	// For any other case, return an UnknownCommand
	return UnknownCommand(commandName)
}

type Command interface {
	Execute() error
}

// ExitCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type ExitCommand string

func (e ExitCommand) Execute() error {
	// If the command has no parameters, exit with code 0
	if e == "" {
		os.Exit(0)
	}
	// Will match "exit 123", not match "exit" or "exit123"
	re := regexp.MustCompile(`^(\d+).*`)
	// Extract the errCode from the command
	errCode, err := strconv.Atoi(string(re.FindSubmatch([]byte(e))[1]))
	// If the conversion fails, consider the command as not an exit command
	if err != nil {
		return errors.New("exit command parameter is not a number")
	}
	os.Exit(errCode)
	return nil
}

// UnknownCommand string contains the command name that is unknown
// It implement the Execute method of the Command interface
type UnknownCommand string

func (u UnknownCommand) Execute() error {
	fmt.Println(u + ": command not found")
	return errors.New("unknown command")
}
