package cmds

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/IdlePhysicist/markup"
)

var xeroxCmd = &cobra.Command{
	Use:   "xerox",
	Short: "Degrade your PDF to look like an old Xerox photocopy",

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pdfFile, err := markup.FindFile(args[0])
		if err != nil {
			log.Fatal(err)
		}

		err = markup.Xerox(pdfFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(xeroxCmd)
}
