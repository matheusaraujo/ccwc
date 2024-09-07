package main

import (
	"os"
	"testing"
)

func runTestFileTest(t *testing.T, countBytes bool, countWords bool, countLines bool, countChars bool, fileName string, content []byte, expectedResult string) {

	if err := os.WriteFile(fileName, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	defer func() {
		if err := os.Remove(fileName); err != nil {
			t.Errorf("Failed to remove test file: %v", err)
		}
	}()

	options := Options{
		CountBytes: countBytes,
		CountWords: countWords,
		CountLines: countLines,
		CountChars: countChars,
	}

	size, err := wc(options, fileName)
	if err != nil {
		t.Fatalf("Error getting file size: %v", err)
	}

	if size != expectedResult {
		t.Errorf("Expected result %s, but got %s", expectedResult, size)
	}
}

func TestCountBytes(t *testing.T) {
	runTestFileTest(t, true, false, false, false, "testfile.txt", []byte("Hello, world!"), "   13 testfile.txt")
}

func TestCountWords(t *testing.T) {
	runTestFileTest(t, false, true, false, false, "testfile.txt", []byte("Hello, happy world!"), "   3 testfile.txt")
}

func TestCountLines(t *testing.T) {
	runTestFileTest(t, false, false, true, false, "testfile.txt", []byte("Hello\nworld\n"), "   2 testfile.txt")
}

func TestCountChars(t *testing.T) {
	runTestFileTest(t, false, false, false, true, "testfile.txt", []byte("Hello\nworld\n"), "   12 testfile.txt")
}

func TestCountEverything(t *testing.T) {
	runTestFileTest(t, true, true, true, true, "testfile.txt", []byte("Hello\nhappy world\n"), "   2 3 18 18 testfile.txt")
}
