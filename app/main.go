package main

import (
	"strings"

	"github.com/chzyer/readline"
	"github.com/codecrafters-io/shell-starter-go/app/command"
)

func main() {
	// Configure readline
	rl, err := readline.New("$ ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	// Set up completion
	rl.Config.AutoComplete = readline.NewPrefixCompleter(command.BuiltinCompletion...)

	// Main shell loop
	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF, readline.ErrInterrupt
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		cmd := command.Factory(line)
		// TODO Execute each command
		cmd[0].Execute()
	}
}
