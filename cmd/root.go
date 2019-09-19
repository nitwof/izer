package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"iconizer/icons"

	"github.com/spf13/cobra"
)

var fontName string
var useColors bool
var checkDirs bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iconizer",
	Short: "Add icons to files by filetypes.",
	Run: func(cmd *cobra.Command, args []string) {
		font := icons.GetFontByName(fontName)
		if font == nil {
			fmt.Printf(
				"Error: Font '%s' is unsupported. Supported fonts: '%s'\n",
				fontName, supportedFonts(),
			)
			cmd.Help() // nolint:errcheck,gosec
			os.Exit(1)
		}

		if len(args) > 0 {
			for _, arg := range args {
				iconize(font, arg)
			}
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				iconize(font, scanner.Text())
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
	rootCmd.PersistentFlags().StringVarP(
		&fontName, "font", "f", "",
		fmt.Sprintf("Font to be used. Supported fonts: %s", supportedFonts()),
	)
	rootCmd.Flags().BoolVarP(
		&useColors, "color", "c", false, "Enable colorful output",
	)
	rootCmd.Flags().BoolVarP(
		&checkDirs, "dir", "d", false,
		"Enable icons for directories (Slows down the process due checking files)",
	)
}

func iconize(font icons.Font, filename string) {
	if useColors {
		fmt.Printf("%s %s\n", getIcon(font, filename).Colored(), filename)
	} else {
		fmt.Printf("%s %s\n", getIcon(font, filename), filename)
	}
}

func getIcon(font icons.Font, filename string) icons.Icon {
	if icon := font.GetIcon(filename); !icon.IsEmpty() {
		return icon
	}

	if checkDirs {
		if stat, err := os.Stat(filename); err == nil && stat.IsDir() {
			return font.DirIcon()
		}
	}

	return font.DefaultIcon()
}

func supportedFonts() string {
	return strings.Join(icons.StringFonts(), ", ")
}
