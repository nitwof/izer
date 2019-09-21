package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

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
