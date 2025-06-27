package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		input, output string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testCases {
		rev, err := Reverse(tc.input)
		if err != nil {
			return
		}
		if rev != tc.output {
			t.Errorf("Reverse(%q) = %q; want %q", tc.input, rev, tc.output)
		}
	}
}

func FuzzReverse(f *testing.F) {
	// testCases := []string{
	// 	"Hello, world",
	// 	" ",
	// 	"!12345",
	// 	"The quick brown fox jumped over the lazy dog",
	// }

	// for _, tc := range testCases {
	// 	f.Add(tc)
	// }

	f.Fuzz(func(t *testing.T, input string) {
		if input == "" {
			return // Skip empty strings
		}
		reversed, err1 := Reverse(input)
		if err1 != nil {
			return
		}
		doubleReversed, err2 := Reverse(reversed)
		if err2 != nil {
			return
		}

		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(input), utf8.RuneCountInString(input), utf8.RuneCountInString(doubleReversed))

		if doubleReversed != input {
			t.Errorf("Reverse(%q) reversed twice does not equal original: got %q, want %q", input, doubleReversed, input)
		}

		if utf8.ValidString(input) && !utf8.ValidString(reversed) {
			t.Errorf("Reverse(%q) produced invalid UTF-8 string: %q", input, reversed)
		}
	})
}
