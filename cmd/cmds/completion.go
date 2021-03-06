package cmds

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Output shell completion code",

	Args: cobra.ExactArgs(1),
	ValidArgs: []string{ "bash", "zsh" },
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case `bash`:
			rootCmd.GenBashCompletion(os.Stdout)
		case `zsh`:
			rootCmd.GenZshCompletion(os.Stdout)
			io.WriteString(os.Stdout, "\ncompdef _markup markup\n")
		default:
			fmt.Printf("Unknown shell: %s", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
