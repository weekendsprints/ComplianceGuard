package cmd

import (
	"os"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	err := versionCommand()
	if err != nil {
		t.Errorf("versionCommand() returned error: %v", err)
	}
}

func TestInfoCommand(t *testing.T) {
	err := infoCommand()
	if err != nil {
		t.Errorf("infoCommand() returned error: %v", err)
	}
}

func TestCheckCommand(t *testing.T) {
	// Save original args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	
	// Test without verbose flag
	os.Args = []string{"complianceguard", "check"}
	err := checkCommand()
	if err != nil {
		t.Errorf("checkCommand() returned error: %v", err)
	}
}
