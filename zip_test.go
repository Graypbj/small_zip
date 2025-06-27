package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestZipAndUnzip(t *testing.T) {
	inputDir := "testdir"
	inputFile := filepath.Join(inputDir, "file.txt")
	archive := "test.zip"
	outputDir := "unzipped"
	buf := make([]byte, 1024)

	// Setup test input
	if err := os.MkdirAll(inputDir, 0755); err != nil {
		t.Fatal(err)
	}
	content := []byte("Zip test content")
	if err := os.WriteFile(inputFile, content, 0644); err != nil {
		t.Fatal(err)
	}

	// Zip
	if err := zipFile(inputDir, archive, buf); err != nil {
		t.Fatalf("zipFolder failed: %v", err)
	}

	// Unzip
	if err := unzipFile(archive, outputDir, buf); err != nil {
		t.Fatalf("unzipFile failed: %v", err)
	}

	// Verify
	resultPath := filepath.Join(outputDir, "file.txt")
	result, err := os.ReadFile(resultPath)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(content, result) {
		t.Errorf("unzipped content mismatch. got: %s", result)
	}

	// Cleanup
	os.RemoveAll(inputDir)
	os.RemoveAll(outputDir)
	os.Remove(archive)
}
