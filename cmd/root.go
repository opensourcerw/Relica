/*
Copyright © 2025 OpenSourceRW <open@opensourcerw.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "relica",
	Short: "relica - AI-powered release note generator",
	Long:  "relica analyzes git commits and generates clean, versioned release notes using AI.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'relica --help' for available commands.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
