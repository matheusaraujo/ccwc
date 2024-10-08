package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetOptions(t *testing.T) {
	tests := []struct {
		name  string
		bytes bool
		words bool
		lines bool
		chars bool
		want  Options
	}{
		{
			name:  "No flags set, defaults to counting bytes, words, and lines",
			bytes: false, words: false, lines: false, chars: false,
			want: Options{CountBytes: true, CountWords: true, CountLines: true, CountChars: false},
		},
		{
			name:  "All flags set to true",
			bytes: true, words: true, lines: true, chars: true,
			want: Options{CountBytes: true, CountWords: true, CountLines: true, CountChars: true},
		},
		{
			name:  "Only CountWords flag set to true",
			bytes: false, words: true, lines: false, chars: false,
			want: Options{CountBytes: false, CountWords: true, CountLines: false, CountChars: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := setOptions(tt.bytes, tt.words, tt.lines, tt.chars)
			assert.Equal(t, tt.want, got)
		})
	}
}
