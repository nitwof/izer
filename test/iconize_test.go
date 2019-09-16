package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	name          string
	args          []string
	inputFixture  string
	goldenFixture string
}{
	{"Nerd", []string{"-f=nerd"}, "nerd.input", "nerd.golden"},
	{"NerdColor", []string{"-f=nerd", "-c"}, "nerd.input", "nerd_color.golden"},
}

func TestIconizeStdin(t *testing.T) {
	binary, err := getBinary()
	require.NoError(t, err, "Cannot get binary path")

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input, err := loadFixture(tt.inputFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.inputFixture)

			golden, err := loadFixture(tt.goldenFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.goldenFixture)

			cmd := exec.Command(binary, tt.args...)
			cmd.Stdin = bytes.NewReader(input)

			result, err := cmd.Output()
			require.NoErrorf(t, err, "Cannot execute binary %s", binary)

			assert.Equal(t, string(golden), string(result))
		})
	}
}

func TestIconizeArgs(t *testing.T) {
	binary, err := getBinary()
	require.NoError(t, err, "Cannot get binary path")

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input, err := loadFixture(tt.inputFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.inputFixture)

			golden, err := loadFixture(tt.goldenFixture)
			assert.NoErrorf(t, err, "Cannot load fixture %s", tt.goldenFixture)

			inputLines := strings.Split(string(input), "\n")
			inputLines = inputLines[0 : len(inputLines)-1]
			args := append(tt.args, inputLines...)

			cmd := exec.Command(binary, args...)

			result, err := cmd.Output()
			require.NoErrorf(t, err, "Cannot execute binary %s", binary)

			assert.Equal(t, string(golden), string(result))
		})
	}
}

func loadFixture(name string) ([]byte, error) {
	file, err := os.Open(filepath.Join("fixtures", name))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getBinary() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	projectPath := filepath.Dir(cwd)
	binaryPath := os.Getenv("ICONIZER")
	if binaryPath == "" {
		binaryPath = "iconizer"
	}

	return filepath.Join(projectPath, binaryPath), nil
}
