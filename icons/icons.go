package icons

var fontsMap = map[string]func(string) Icon{
	"nerd": GetNerdIcon,
}
var supportedFonts = []string{
	"nerd",
}

// Icon represents icon struct
type Icon struct {
	Symbol string
	Color  uint8
}

// SupportedFonts returns list of supported fonts
func SupportedFonts() []string {
	return supportedFonts
}

// GetIconFunc returns GetIconFunc by font
func GetIconFunc(font string) func(string) Icon {
	if fn, ok := fontsMap[font]; ok {
		return fn
	}
	return nil
}
