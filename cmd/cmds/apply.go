package cmds

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/IdlePhysicist/markup"
)

var (
	applyCmd = &cobra.Command{
		Use:   "apply",
		Short: "Apply a LaTeX template to a Markdown file and render to PDF",

		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			mdFile, err := markup.OpenFile(args[0])
			if err != nil {
				log.Fatal(err)
			}

			template, err := markup.FindTemplate(templateName)
			if err != nil {
				log.Fatal(err)
			}

			err = markup.Markup(mdFile, template)
			if err != nil {
				log.Fatal(err)
			}

			if xeroxFlg {
				err = markup.Xerox(mdFile)
				if err != nil {
					log.Fatal(err)
				}
			}

		},
	}

	xeroxFlg     bool
	templateName string
)

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flags().BoolVarP(&xeroxFlg, "xerox", "x", false, "Xerox the output")
	applyCmd.Flags().StringVarP(&templateName, "template", "t", "", "Template name")
	applyCmd.MarkFlagRequired("template")
}
