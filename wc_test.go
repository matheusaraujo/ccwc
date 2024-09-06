package main

import (
	"os"
	"testing"
)

func TestWc(t *testing.T) {
	testFile := "testfile.txt"
	content := []byte("Hello, world!")
	err := os.WriteFile(testFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	defer os.Remove(testFile)

	size, err := wc(testFile)
	if err != nil {
		t.Fatalf("Error getting file size: %v", err)
	}

	expectedSize := "13 testfile.txt"

	if size != expectedSize {
		t.Errorf("Expected file size %s, but got %s", expectedSize, size)
	}
}
