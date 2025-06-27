package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGzipAndGunzip(t *testing.T) {
	input := "test_input.txt"
	compressed := "test_output.gz"
	output := "test_output.txt"
	content := []byte("This is a test.")

	if err := os.WriteFile(input, content, 0644); err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1024)

	if err := gzipFile(input, compressed, buf); err != nil {
		t.Fatalf("gzipFile failed: %v", err)
	}

	if err := gunzipFile(compressed, output, buf); err != nil {
		t.Fatalf("gunzipFile failed: %v", err)
	}

	result, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(content, result) {
		t.Errorf("gunzip result does not match original. got: %s", result)
	}

	os.Remove(input)
	os.Remove(compressed)
	os.Remove(output)
}
