package main

import (
	"testing"
)

func TestCommandFactoryForExitCommand(t *testing.T) {
	// Test with an exit command
	command := "exit 123"
	want := ExitCommand("123")

	if got := commandFactory(command); got != want {
		t.Errorf("commandFactory(%q) = %q; want %q", command, got, want)
	}

	command = "exit"
	want = ExitCommand("")
	if got := commandFactory(command); got != want {
		t.Errorf("commandFactory(%q) = %q; want %q", command, got, want)
	}

	command = "exit123"
	want2 := UnknownCommand("exit123")
	if got := commandFactory(command); got != want2 {
		t.Errorf("commandFactory(%q) = %q; want %q", command, got, want2)
	}
}

func TestCommandFactoryForUnknownCommand(t *testing.T) {
	// Test with an unknown command
	command := "unknown_command"
	want := UnknownCommand("unknown_command")

	if got := commandFactory(command); got != want {
		t.Errorf("commandFactory(%q) = %q; want %q", command, got, want)
	}
}
