package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion bash",
	Short: "Generates completion scripts for the given shell",
	Long: `Generates completion scripts for the given shell

To load bash completion run

. <(izer completion bash)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(izer completion bash)
	`,
	ValidArgs: []string{"bash"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			if err := rootCmd.GenBashCompletion(os.Stdout); err != nil {
				fmt.Printf("Cannot generate bash completion: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
