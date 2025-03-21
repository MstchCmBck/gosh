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

func TestIsExit(t *testing.T) {
	// Test with a command that is not an exit command
	command := "exit"
	want := false

	if got, _ := isExit(command); got != want {
		t.Errorf("isExit(%q) = %t; want %t", command, got, want)
	}

	// Test with a command that is an exit command
	command = "exit123"
	want = false

	if got, _ := isExit(command); got != want {
		t.Errorf("isExit(%q) = %t; want %t", command, got, want)
	}

	// Test with a command that is an exit command
	command = "exit 0"
	want = true
	wantIndex := 0

	if got, index := isExit(command); got != want && index != wantIndex {
		t.Errorf("isExit(%q) = %t; want %t", command, got, want)
		t.Errorf("isExit(%q) index = %d; wantIndex %d", command, index, wantIndex)
	}

	// Test with a command that is an exit command
	command = "exit 123"
	want = true
	wantIndex = 123

	if got, index := isExit(command); got != want && index != wantIndex {
		t.Errorf("isExit(%q) = %t; want %t", command, got, want)
		t.Errorf("isExit(%q) index = %d; wantIndex %d", command, index, wantIndex)
	}
}
