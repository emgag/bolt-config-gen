package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/emgag/bolt-config-gen/internal/lib/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bolt-config-gen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("bolt-config-gen %s -- %s\n", version.Version, version.Commit)
	},
}
