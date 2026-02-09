package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display system information",
	Long:  `Display detailed information about the system including OS, architecture, CPUs, and Go version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("System Information:")
		fmt.Printf("  OS:           %s\n", runtime.GOOS)
		fmt.Printf("  Architecture: %s\n", runtime.GOARCH)
		fmt.Printf("  CPUs:         %d\n", runtime.NumCPU())
		fmt.Printf("  Go Version:   %s\n", runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
