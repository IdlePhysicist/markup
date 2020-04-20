package cmd

import (

	"github.com/spf13/cobra"

	"github.com/IdlePhysicist/markup/toolchest"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A list the known templates",

	Run: func(cmd *cobra.Command, args []string) {
		_ = toolchest.FindAllTemplates()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
