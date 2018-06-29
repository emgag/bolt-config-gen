package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bolt-config-gen",
	Short: "Generate Bolt CMS configs from template files",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.Execute()
}

