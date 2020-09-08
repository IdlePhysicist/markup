package cmds

import (

	"github.com/spf13/cobra"

	"github.com/IdlePhysicist/markup"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A list the known templates",

	Run: func(cmd *cobra.Command, args []string) {
		_ = markup.FindAllTemplates()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
