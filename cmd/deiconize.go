package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var deiconizeCmd = &cobra.Command{
	Use:   "deiconize",
	Short: "Remove icons from filenames (undo iconize)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, arg := range args {
				deiconize(arg)
			}
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				deiconize(scanner.Text())
			}

			if scanner.Err() != nil {
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deiconizeCmd)
}

func deiconize(filename string) {
	cutIndex := strings.IndexByte(filename, ' ')
	if cutIndex == -1 {
		fmt.Println(filename)
	} else {
		fmt.Println(filename[cutIndex+1:])
	}
}
