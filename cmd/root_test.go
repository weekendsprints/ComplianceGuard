package cmd

import (
	"bytes"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	// Capture output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	
	// Set args and execute
	rootCmd.SetArgs([]string{"version"})
	
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("version command returned error: %v", err)
	}
	
	// Reset for other tests
	rootCmd.SetArgs([]string{})
}

func TestInfoCommand(t *testing.T) {
	// Capture output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	
	// Set args and execute
	rootCmd.SetArgs([]string{"info"})
	
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("info command returned error: %v", err)
	}
	
	// Reset for other tests
	rootCmd.SetArgs([]string{})
}

func TestCheckCommand(t *testing.T) {
	// Capture output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	
	// Set args and execute (without GitLab flags, so it skips GitLab verification)
	rootCmd.SetArgs([]string{"check"})
	
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("check command returned error: %v", err)
	}
	
	// Reset for other tests
	rootCmd.SetArgs([]string{})
}

func TestCheckCommandWithVerbose(t *testing.T) {
	// Capture output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	
	// Set args with verbose flag
	rootCmd.SetArgs([]string{"check", "--verbose"})
	
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("check command with verbose returned error: %v", err)
	}
	
	// Reset for other tests
	rootCmd.SetArgs([]string{})
}
