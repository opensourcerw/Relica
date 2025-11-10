/*
Copyright © 2025 OpenSourceRW <open@opensourcerw.com>
*/
package cmd

import (
	"os"
	"fmt"
	"time"
	
	"github.com/spf13/cobra"
	"github.com/opensourcerw/relica/internal/git"
	"github.com/opensourcerw/relica/internal/config"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Relica in the current repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Step 1: Check for Git repo
		if !git.IsGitRepo(".") {
			return fmt.Errorf("no git repository found in the current directory")
		}

		// Step 2: Create .relica/ folder & config file
		if err := os.MkdirAll(".relica", 0755); err != nil {
			return fmt.Errorf("failed to create .relica config directory: %w", err)
		}
		if err := os.MkdirAll("relica", 0755); err != nil {
			return fmt.Errorf("failed to create relica release directory: %w", err)
		}

		if config.Exists() {
			return fmt.Errorf("Relica already initialized in this repository")
		}

		cfg := config.Config{
			Initialized: true,
			RepoPath:    ".",
			CreatedAt:   time.Now(),
			NotesDir:    "relica/",
			Version:     "1.0.0",
		}

		if err := config.Save(cfg); err != nil {
			return fmt.Errorf("failed to save Relica config: %w", err)
		}

		fmt.Println("✅ Relica initialized successfully in this repo.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}