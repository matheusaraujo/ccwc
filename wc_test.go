package main

import (
	"testing"
)

func TestWc(t *testing.T) {
	expected := "Hello, World!"
	result := wc()

	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}
