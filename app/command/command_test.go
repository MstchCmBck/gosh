package command

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestCommandFactoryForExitCommand(t *testing.T) {
	// Test with an exit command
	command := "exit 123"
	want := ExitCommand("123")

	if got := Factory(command); got != want {
		t.Errorf("commandFactory(%q) = %q; want %q", command, got, want)
	}

	command = "exit"
	want = ExitCommand("")
	if got := Factory(command); got != want {
		t.Errorf("commandFactory(%q) = %q; want %q", command, got, want)
	}

	command = "exit123"
	want2 := UnknownCommand("exit123")
	if got := Factory(command); got != want2 {
		t.Errorf("commandFactory(%q) = %q; want %q", command, got, want2)
	}
}

func TestCommandFactoryForUnknownCommand(t *testing.T) {
	// Test with an unknown command
	command := "unknown_command"
	want := UnknownCommand("unknown_command")

	if got := Factory(command); got != want {
		t.Errorf("commandFactory(%q) = %q; want %q", command, got, want)
	}
}

func TestUnknownCommand(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run command
	cmd := UnknownCommand("testcmd")
	err := cmd.Execute()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := strings.TrimSpace(buf.String())

	// Assert expectations
	if err == nil {
		t.Error("Expected error but got none")
	}
	if err.Error() != "unknown command" {
		t.Errorf("Expected error 'unknown command' but got '%v'", err)
	}
	if output != "testcmd: command not found" {
		t.Errorf("Expected output 'testcmd: command not found' but got '%v'", output)
	}
}
