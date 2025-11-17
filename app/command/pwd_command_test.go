// pwd_command_test.go
package command

import (
	"bytes"
	"os"
	"testing"
)

func TestPwdCommand_Execute(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "gosh_pwd_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test

	// Change the current working directory to the temp directory
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}
	defer func() {
		os.Chdir(oldWd)
	}()

	// Create a pwdcommand instance
	var out bytes.Buffer
	params := parameters{
		stdout: &out,
	}
	cmd := pwdcommand(params)

	// Execute the command
	err = cmd.execute()
	if err != nil {
		t.Errorf("execute() returned an error: %v", err)
	}

	// Get the expected output
	expectedWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	expectedOutput := expectedWd + "\n"

	// Assert the output
	actualOutput := out.String()
	if actualOutput != expectedOutput {
		t.Errorf("execute() output:\nExpected: %q\nGot:      %q", expectedOutput, actualOutput)
	}
}
