package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version information
	version = "1.0.0"
	
	// Global flags
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "complianceguard",
	Short: "ComplianceGuard - Cross-Platform CLI Application",
	Long: `ComplianceGuard is a command-line tool that helps you run compliance checks
and verify configurations across different platforms.`,
}

// Execute runs the CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
