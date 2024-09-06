package main

import (
	"os"
	"testing"
)

func runTestFileTest(t *testing.T, countBytes bool, fileName string, content []byte, expectedResult string) {

	if err := os.WriteFile(fileName, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	defer func() {
		if err := os.Remove(fileName); err != nil {
			t.Errorf("Failed to remove test file: %v", err)
		}
	}()

	size, err := wc(countBytes, fileName)
	if err != nil {
		t.Fatalf("Error getting file size: %v", err)
	}

	if size != expectedResult {
		t.Errorf("Expected result %s, but got %s", expectedResult, size)
	}
}

func TestCountBytes(t *testing.T) {
	runTestFileTest(t, true, "testfile.txt", []byte("Hello, world!"), "13 testfile.txt")
}
