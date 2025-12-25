package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Global flags
	humanOutput bool
	teamID      string
	projectID   string
)

// NewRootCmd creates the root command for the Linear CLI
func NewRootCmd(version, commit, date string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "linear",
		Short: "Linear CLI - AI-optimized project management",
		Long: `Linear Agent CLI - A command-line interface for Linear project management.

Designed for AI agent consumption with JSON-first output.
Use --human flag for human-readable output.

Configuration:
  linear config setup    Interactive setup wizard
  linear config set      Set configuration values
  linear config get      Get configuration values

Quick Start:
  linear issue list      List issues in current project
  linear project list    List all projects
  linear document list   List documents`,
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Load configuration before each command
			// This will be implemented in config package
		},
	}

	// Global flags
	rootCmd.PersistentFlags().BoolVar(&humanOutput, "human", false, "Output in human-readable format (default: JSON)")
	rootCmd.PersistentFlags().StringVar(&teamID, "team", "", "Team ID or key (overrides config)")
	rootCmd.PersistentFlags().StringVar(&projectID, "project", "", "Project ID (overrides VCS detection)")

	// Add command groups
	rootCmd.AddCommand(NewAuthCmd())
	rootCmd.AddCommand(NewIssueCmd())
	rootCmd.AddCommand(NewProjectCmd())
	rootCmd.AddCommand(NewDocumentCmd())
	rootCmd.AddCommand(NewLabelCmd())
	rootCmd.AddCommand(NewWorkflowCmd())
	rootCmd.AddCommand(NewStatusCmd())
	rootCmd.AddCommand(NewUserCmd())
	rootCmd.AddCommand(NewTeamCmd())
	rootCmd.AddCommand(NewInitiativeCmd())
	rootCmd.AddCommand(NewConfigCmd())
	rootCmd.AddCommand(NewWhoamiCmd())

	return rootCmd
}

// OutputJSON outputs data as JSON (default mode)
func OutputJSON(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// OutputHuman outputs data in human-readable format
func OutputHuman(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// IsHumanOutput returns whether human output mode is enabled
func IsHumanOutput() bool {
	return humanOutput
}

// GetTeamID returns the team ID from flag or config
func GetTeamID() string {
	return teamID
}

// GetProjectID returns the project ID from flag or VCS detection
func GetProjectID() string {
	return projectID
}
