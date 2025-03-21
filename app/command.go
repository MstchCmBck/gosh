package main

func unknownCommand(command string) string {
	return command[:len(command)-1] + ": command not found"
}
