package main

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkDeiconize(b *testing.B) {
	binary, err := getBinary()
	require.NoError(b, err, "Cannot get binary path")

	inputData, err := loadFixture("deiconize_bench.input")
	require.NoErrorf(b, err, "Cannot load fixture deiconize_bench.input")

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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmd := exec.Command(binary, "deiconize")
		cmd.Stdin = bytes.NewReader(input)

		if output, err := cmd.CombinedOutput(); err != nil {
			b.Errorf("Error on execution command %s: %v", binary, err)
			b.Errorf("%s", output)
		}
	}
}
