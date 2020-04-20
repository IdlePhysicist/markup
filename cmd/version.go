package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s (%s)\n", Version, Commit)
		},
	}

	Commit  string
	Version string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
