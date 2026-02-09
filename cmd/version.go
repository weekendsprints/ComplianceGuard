package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Display version information including Go version and build details.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ComplianceGuard v%s\n", version)
		fmt.Printf("Built with %s for %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
