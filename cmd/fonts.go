package cmd

import (
	"fmt"

	"iconizer/icons"

	"github.com/spf13/cobra"
)

var fontsCmd = &cobra.Command{
	Use:   "fonts",
	Short: "List of supported fonts",
	Run: func(cmd *cobra.Command, args []string) {
		for _, font := range icons.Fonts() {
			fmt.Println(font)
		}
	},
}

func init() {
	rootCmd.AddCommand(fontsCmd)
}
