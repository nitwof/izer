package icons

import (
	"fmt"
)

// Font interface represents abstract font
type Font interface {
	fmt.Stringer
	GetIcon(filename string) Icon
	DefaultIcon() Icon
	DirIcon() Icon
}

// GetFontByName returns font by it string representation
func GetFontByName(name string) Font {
	for _, font := range Fonts() {
		if font.String() == name {
			return font
		}
	}

	return nil
}

// Fonts returns list of supported fonts
func Fonts() []Font {
	return []Font{
		NerdFont{},
	}
}

// StringFonts returns list of string representations of supported fonts
func StringFonts() []string {
	fonts := Fonts()
	strFonts := make([]string, len(fonts))

	for i, font := range fonts {
		strFonts[i] = font.String()
	}

	return strFonts
}
