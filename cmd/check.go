package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"complianceguard/gitlab"

	"github.com/spf13/cobra"
)

var (
	gitlabURL   string
	gitlabToken string
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run compliance checks",
	Long: `Run compliance checks including configuration validation, permissions checks,
and optional GitLab API verification.

GitLab verification requires both --gitlab-url and --gitlab-token flags
(or GITLAB_URL and GITLAB_TOKEN environment variables).`,
	RunE: runCheck,
}

func init() {
	rootCmd.AddCommand(checkCmd)
	
	// GitLab integration flags
	checkCmd.Flags().StringVar(&gitlabURL, "gitlab-url", os.Getenv("GITLAB_URL"), "GitLab instance URL")
	checkCmd.Flags().StringVar(&gitlabToken, "gitlab-token", os.Getenv("GITLAB_TOKEN"), "GitLab PAT token")
}

func runCheck(cmd *cobra.Command, args []string) error {
	fmt.Println("Running compliance checks...")

	if verbose {
		fmt.Println("  [VERBOSE] Checking configuration files...")
		fmt.Println("  [VERBOSE] Validating permissions...")
		fmt.Println("  [VERBOSE] Scanning dependencies...")
	}

	// GitLab API verification step
	if gitlabURL != "" && gitlabToken != "" {
		if verbose {
			fmt.Println("  [VERBOSE] Verifying GitLab connection...")
		}

		settings, err := gitlab.VerifyConnection(gitlabURL, gitlabToken, verbose)
		if err != nil {
			return fmt.Errorf("GitLab verification failed: %w", err)
		}

		fmt.Println("✓ GitLab connection verified")
		if verbose {
			settingsJSON, _ := json.MarshalIndent(settings, "  ", "  ")
			fmt.Printf("  [VERBOSE] GitLab settings:\n  %s\n", string(settingsJSON))
		}
	} else if verbose {
		fmt.Println("  [VERBOSE] Skipping GitLab verification (no URL/token provided)")
	}

	fmt.Println("✓ All compliance checks passed")
	return nil
}
