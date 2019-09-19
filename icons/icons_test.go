package icons

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColored(t *testing.T) {
	icon := Icon{"O", 7}
	expected := fmt.Sprintf("\x1b[37m%s\x1b[0m", icon.Symbol)

	result := icon.Colored().String()
	assert.Equal(t, expected, result)
}

func TestIsEmpty(t *testing.T) {
	t.Run("IconEmpty", func(t *testing.T) {
		icon := Icon{}
		assert.True(t, icon.IsEmpty())
	})
	t.Run("IconNotEmpty", func(t *testing.T) {
		icon := Icon{"0", 7}
		assert.False(t, icon.IsEmpty())
	})
}

func TestString(t *testing.T) {
	icon := Icon{"O", 7}
	assert.Equal(t, icon.Symbol, icon.String())
}
