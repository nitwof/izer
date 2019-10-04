package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion bash|zsh",
	Short: "Generates completion scripts for the given shell",
	Long: `Generates completion scripts for the given shell

To load bash completion run:
. <(izer completion bash)

To load zsh completion run:
. <(izer completion zsh)

To configure your shell to load completions for each session
add this to your shell rc file:

# ~/.bashrc or ~/.profile
. <(izer completion bash)

# ~/.zshrc
. <(izer completion zsh)
	`,
	ValidArgs: []string{"bash", "zsh"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			if err := rootCmd.GenBashCompletion(os.Stdout); err != nil {
				fmt.Printf("Cannot generate bash completion: %v\n", err)
				os.Exit(1)
			}
		case "zsh":
			if err := rootCmd.GenZshCompletion(os.Stdout); err != nil {
				fmt.Printf("Cannot generate zsh completion: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
