package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var deiconizeTests = []struct {
	name          string
	args          []string
	inputFixture  string
	goldenFixture string
}{
	{"Nerd", []string{}, "nerd/deiconize.input", "nerd/deiconize.golden"},
	{"NerdColor", []string{}, "nerd/deiconize_color.input", "nerd/deiconize.golden"},
	{"NerdDir", []string{}, "nerd/deiconize_dir.input", "nerd/deiconize.golden"},
	{"NerdDirColor", []string{}, "nerd/deiconize_dir_color.input", "nerd/deiconize.golden"},
}

func TestDeiconizeStdin(t *testing.T) {
	binary, err := getBinary()
	require.NoError(t, err, "Cannot get binary path")

	for _, tt := range deiconizeTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input, err := loadFixture(tt.inputFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.inputFixture)

			golden, err := loadFixture(tt.goldenFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.goldenFixture)

			cmdArgs := append([]string{"deiconize"}, tt.args...)
			cmd := exec.Command(binary, cmdArgs...)
			cmd.Stdin = bytes.NewReader(input)

			result, err := cmd.Output()
			require.NoErrorf(t, err, "Cannot execute binary %s", binary)

			assert.Equal(t, string(golden), string(result))
		})
	}
}

func TestDeiconizeArgs(t *testing.T) {
	binary, err := getBinary()
	require.NoError(t, err, "Cannot get binary path")

	for _, tt := range deiconizeTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input, err := loadFixture(tt.inputFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.inputFixture)

			golden, err := loadFixture(tt.goldenFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.goldenFixture)

			inputLines := strings.Split(string(input), "\n")
			inputLines = inputLines[0 : len(inputLines)-1]

			cmdArgs := append([]string{"deiconize"}, append(tt.args, inputLines...)...)
			cmd := exec.Command(binary, cmdArgs...)

			result, err := cmd.Output()
			require.NoErrorf(t, err, "Cannot execute binary %s", binary)

			assert.Equal(t, string(golden), string(result))
		})
	}
}
