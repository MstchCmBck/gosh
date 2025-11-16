package main

import (
	"fmt"
	"strings"
	"sync"

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

		commands := command.Factory(line)

		var wg sync.WaitGroup

		for _, cmd := range commands {
			wg.Add(1)
			go func(c command.Command) {
				defer wg.Done()
				err := c.Execute()
				if err != nil {
					fmt.Printf("Command failed: %v\n", err)
				}
			}(cmd)
		}

		wg.Wait()
	}
}
