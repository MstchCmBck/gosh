package command

import (
	"fmt"
	"os"
	"strings"
)

// TypeCommand string contains parameters send with the exit command
// It implement the Execute method of the Command interface
type TypeCommand string

var execMap = make(map[string]string)

func init() {
	builtinCommands["type"] = func(params string) Command {
		return TypeCommand(params)
	}
	execMap = getExecMap()
}

func (t TypeCommand) Execute() error {
	_, find := builtinCommands[string(t)]
	if find {
		fmt.Println(t + " is a shell builtin")
		return nil
	}

	path, exists := execMap[string(t)]
	if !exists {
		fmt.Println(t + ": not found")
		return nil
	}
	fmt.Println(string(t) + " is " + path + "/" + string(t))
	return nil
}

func getExecMap() map[string]string {
	path, exists := os.LookupEnv("PATH")
	if !exists {
		return nil
	}
	splitted := strings.Split(path, ":")

	var emap = make(map[string]string)

	for _, dir := range splitted {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, file := range files {
			emap[file.Name()] = dir
		}
	}

	return emap
}
