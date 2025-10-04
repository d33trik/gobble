package main_test

import (
	"strings"
	"testing"

	gobble "github.com/d33trik/gobble"
)

func TestCountLines(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"one new line": {
			input: "one two three four five\n",
			want:  1,
		},
		"empty input": {
			input: "",
			want:  0,
		},
		"no new lines": {
			input: "one two three four five six",
			want:  0,
		},
		"no new lines at end": {
			input: "one two three four five\nsix",
			want:  1,
		},
		"multiple lines with words": {
			input: "one\ntow\nthree\nfour\nfive\n",
			want:  5,
		},
		"multiple lines without words": {
			input: "\n\n\n\n\n",
			want:  5,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			got := gobble.CountLines(r)
			if got != tc.want {
				t.Logf("got: %d, want: %d", got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"five words": {
			input: "one two three four five",
			want:  5,
		},
		"empty input": {
			input: "",
			want:  0,
		},
		"single space": {
			input: " ",
			want:  0,
		},
		"multiple spaces": {
			input: "one two three  four five six",
			want:  6,
		},
		"start multiple spaces": {
			input: "  one two three four five six",
			want:  6,
		},
		"end multiple spaces": {
			input: "one two three four five six",
			want:  6,
		},
		"new line": {
			input: "one two three\nfour five six",
			want:  6,
		},
		"tab character": {
			input: "one two three\tfour five six",
			want:  6,
		},
		"utf8 spaces": {
			input: "one two three four five six",
			want:  6,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			got := gobble.CountWords(r)
			if got != tc.want {
				t.Logf("got: %d, want: %d", got, tc.want)
				t.Fail()
			}
		})
	}
}
