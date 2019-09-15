package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"iconizer/icons"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	shiftFg = 16
	flagFg  = (1 << 14)
)

var font string
var color bool
var supportedFonts string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iconizer",
	Short: "Add icons to files by filetypes.",
	Run: func(cmd *cobra.Command, args []string) {
		getIconFunc := icons.GetIconFunc(font)
		if getIconFunc == nil {
			fmt.Printf(
				"Error: Font %s is unsupported. Supported fonts: %s\n",
				font, supportedFonts,
			)
			cmd.Help() // nolint:errcheck
			os.Exit(1)
		}

		if len(args) > 0 {
			for _, arg := range args {
				printIconFilename(getIconFunc(arg), arg)
			}
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				printIconFilename(getIconFunc(scanner.Text()), scanner.Text())
			}

			if scanner.Err() != nil {
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// main.main() calls this. It needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	supportedFonts = strings.Join(icons.SupportedFonts(), ", ")

	rootCmd.PersistentFlags().StringVarP(
		&font, "font", "f", "",
		fmt.Sprintf("Font to be used. Supported fonts: %s", supportedFonts),
	)
	rootCmd.Flags().BoolVarP(
		&color, "color", "c", false, "Enable colorful output",
	)
}

func printIconFilename(icon icons.Icon, filename string) {
	if color {
		color := (aurora.Color(icon.Color) << shiftFg) | flagFg
		fmt.Printf("%s %s\n", aurora.Colorize(icon.Symbol, color), filename)
	} else {
		fmt.Printf("%s %s\n", icon.Symbol, filename)
	}
}
