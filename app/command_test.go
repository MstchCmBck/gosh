package main

import (
	"testing"
)

// TestUnknownCommand tests the unknownCommand function
func TestUnknownCommand(t *testing.T) {
	// As the user enter its command with the enter key, the gathered command will have a trailing newline character
	command := "unknown_command\n"
	want := "unknown_command: command not found"

	if want != unknownCommand(command) {
		t.Errorf("unknownCommand(%q) = %q; want %q", command, unknownCommand(command), want)
	}
}	
