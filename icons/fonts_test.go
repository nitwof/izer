package icons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFontByName(t *testing.T) {
	tests := []struct {
		name     string
		fontName string
		result   Font
	}{
		{"Nerd", "nerd", NerdFont{}},
		{"Unknown", "unknown", nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := GetFontByName(tt.fontName)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestFonts(t *testing.T) {
	expected := []Font{
		NerdFont{},
	}

	result := Fonts()
	assert.Equal(t, expected, result)
}

func TestStringFonts(t *testing.T) {
	expected := []string{
		"nerd",
	}

	result := StringFonts()
	assert.Equal(t, expected, result)
}
