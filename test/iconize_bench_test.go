package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkIconize(b *testing.B) {
	binary, err := getBinary()
	require.NoError(b, err, "Cannot get binary path")

	projectPath, err := filepath.Abs("..")
	require.NoError(b, err, "Cannot get project path")

	inputData, err := loadFixture("iconize_bench.input")
	require.NoErrorf(b, err, "Cannot load fixture iconize_bench.input")

	// Data multiplicator
	mul := 1000
	// Generate input data
	input := make([]byte, len(inputData)*mul)
	for i := 0; i < mul; i++ {
		for j := range inputData {
			input[(i*len(inputData))+j] = inputData[j]
		}
	}

	b.Logf(
		"Input size: %d KB - %d lines",
		len(input)/1024.0, bytes.Count(input, []byte("\n")),
	)

	for _, tt := range iconizeTests {
		tt := tt
		b.Run(tt.name, func(b *testing.B) {
			cmdArgs := append([]string{"iconize"}, tt.args...)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cmd := exec.Command(binary, cmdArgs...)
				cmd.Stdin = bytes.NewReader(input)
				cmd.Dir = projectPath

				if output, err := cmd.CombinedOutput(); err != nil {
					b.Errorf("Error on execution command %s: %v", binary, err)
					b.Errorf("%s", output)
				}
			}
		})
	}
}
