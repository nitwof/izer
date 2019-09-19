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
	{"Nerd", []string{"-f=nerd"}, "nerd/iconize.input", "nerd/iconize.golden"},
	{"NerdColor", []string{"-f=nerd", "-c"}, "nerd/iconize.input", "nerd/iconize_color.golden"},
	{"NerdDir", []string{"-f=nerd", "-d"}, "nerd/iconize.input", "nerd/iconize_dir.golden"},
	{"NerdDirColor", []string{"-f=nerd", "-c", "-d"}, "nerd/iconize.input", "nerd/iconize_dir_color.golden"},
}

func TestIconizeStdin(t *testing.T) {
	binary, err := getBinary()
	require.NoError(t, err, "Cannot get binary path")

	projectPath, err := filepath.Abs("..")
	require.NoError(t, err, "Cannot get project path")

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input, err := loadFixture(tt.inputFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.inputFixture)

			golden, err := loadFixture(tt.goldenFixture)
			require.NoErrorf(t, err, "Cannot load fixture %s", tt.goldenFixture)

			cmd := exec.Command(binary, tt.args...)
			cmd.Stdin = bytes.NewReader(input)
			cmd.Dir = projectPath

			result, err := cmd.Output()
			require.NoErrorf(t, err, "Cannot execute binary %s", binary)

			assert.Equal(t, string(golden), string(result))
		})
	}
}

func TestIconizeArgs(t *testing.T) {
	binary, err := getBinary()
	require.NoError(t, err, "Cannot get binary path")

	projectPath, err := filepath.Abs("..")
	require.NoError(t, err, "Cannot get project path")

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
			cmd.Dir = projectPath

			result, err := cmd.Output()
			require.NoErrorf(t, err, "Cannot execute binary %s", binary)

			assert.Equal(t, string(golden), string(result))
		})
	}
}

func loadFixture(name string) ([]byte, error) {
	file, err := os.Open(filepath.Join("fixtures", filepath.Clean(name)))
	if err != nil {
		return nil, err
	}
	defer file.Close() // nolint:errcheck

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
