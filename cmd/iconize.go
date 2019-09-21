package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NightWolf007/izer/icons"

	"github.com/spf13/cobra"
)

var iconizeFontName string
var iconizeUseColors bool
var iconizeCheckDirs bool

var iconizeCmd = &cobra.Command{
	Use:   "iconize",
	Short: "Add icons to filenames",
	Run: func(cmd *cobra.Command, args []string) {
		font := icons.GetFontByName(iconizeFontName)
		if font == nil {
			fmt.Printf(
				"Error: Font '%s' is unsupported. See: izer fonts",
				iconizeFontName,
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

func init() {
	iconizeCmd.Flags().StringVarP(
		&iconizeFontName, "font", "f", "",
		fmt.Sprintf("Font to be used. See: izer fonts"),
	)
	iconizeCmd.Flags().BoolVarP(
		&iconizeUseColors, "color", "c", false, "Enable colorful output",
	)
	iconizeCmd.Flags().BoolVarP(
		&iconizeCheckDirs, "dir", "d", false,
		"Enable icons for directories (Slows down the process due checking files)",
	)

	rootCmd.AddCommand(iconizeCmd)
}

func iconize(font icons.Font, filename string) {
	if iconizeUseColors {
		fmt.Printf("%s %s\n", getIcon(font, filename).Colored(), filename)
	} else {
		fmt.Printf("%s %s\n", getIcon(font, filename), filename)
	}
}

func getIcon(font icons.Font, filename string) icons.Icon {
	if icon := font.GetIcon(filename); !icon.IsEmpty() {
		return icon
	}

	if iconizeCheckDirs {
		if stat, err := os.Stat(filename); err == nil && stat.IsDir() {
			return font.DirIcon()
		}
	}

	return font.DefaultIcon()
}
