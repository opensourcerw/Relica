package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/opensourcerw/relica/pkg/version"
)

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print Relica version information",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Relica %s\nCommit: %s\nBuilt: %s\n",
            version.Version, version.Commit, version.BuildDate)
    },
}

func init() {
    rootCmd.AddCommand(versionCmd)
}