package icons

import (
	"github.com/logrusorgru/aurora"
)

// Icon represents icon struct
type Icon struct {
	Symbol string
	Color  uint8
}

// Colored returns colorful ANSI representation of icon
func (icon Icon) Colored() aurora.Value {
	color := (aurora.Color(icon.Color) << 16) | (1 << 14)
	return aurora.Colorize(icon.Symbol, color)
}

// IsEmpty checks if icon is empty
func (icon Icon) IsEmpty() bool {
	return icon.Symbol == ""
}

// String returns string representation of icon
func (icon Icon) String() string {
	return icon.Symbol
}
