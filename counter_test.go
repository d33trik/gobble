package main_test

import (
	"bytes"
	"strings"
	"testing"

	gobble "github.com/d33trik/gobble"
)

func TestCountLines(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"empty input": {
			input: "",
			want:  0,
		},
		"only words": {
			input: "one two three four five six",
			want:  0,
		},
		"only new lines": {
			input: "\n\n\n\n\n",
			want:  5,
		},
		"new line": {
			input: "one two three\nfour five six",
			want:  1,
		},
		"end new line": {
			input: "one two three four five six\n",
			want:  1,
		},
		"multiple new lines": {
			input: "one\ntow\nthree\nfour\nfive\nsix\n",
			want:  6,
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
		"empty input": {
			input: "",
			want:  0,
		},
		"only words": {
			input: "one two three four five",
			want:  5,
		},
		"only spaces": {
			input: "       ",
			want:  0,
		},
		"multiple spaces between words": {
			input: "one two three  four five six",
			want:  6,
		},
		"start multiple spaces": {
			input: "  one two three four five six",
			want:  6,
		},
		"end multiple spaces": {
			input: "one two three four five six  ",
			want:  6,
		},
		"utf8 spaces": {
			input: "one two three four five six",
			want:  6,
		},
		"unicode characters": {
			input: "Ђ ʩ",
			want:  2,
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

func TestCountBytes(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"empty input": {
			input: "",
			want:  0,
		},
		"only words": {
			input: "one two three four five six",
			want:  27,
		},
		"only new lines": {
			input: "\n\n\n\n\n",
			want:  5,
		},
		"multiple new lines": {
			input: "one\ntwo\nthree\nfour\nfive\nsix",
			want:  27,
		},
		"only spaces": {
			input: "       ",
			want:  7,
		},
		"multiple spaces between words": {
			input: "one two three  four five six",
			want:  28,
		},
		"utf8 spaces": {
			input: "one two three four five six",
			want:  37,
		},
		"unicode characters": {
			input: "Ђ ʩ",
			want:  5,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			got := gobble.CountBytes(r)
			if got != tc.want {
				t.Logf("got: %d, want: %d", got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := map[string]struct {
		input string
		want  gobble.Stats
	}{
		"empty input": {
			input: "",
			want: gobble.Stats{
				Lines: 0,
				Words: 0,
				Bytes: 0,
			},
		},
		"only words": {
			input: "one two three four five",
			want: gobble.Stats{
				Lines: 0,
				Words: 5,
				Bytes: 23,
			},
		},
		"only new lines": {
			input: "\n\n\n\n\n",
			want: gobble.Stats{
				Lines: 5,
				Words: 0,
				Bytes: 5,
			},
		},
		"multiple new lines": {
			input: "one\ntwo\nthree\nfour\nfive\n",
			want: gobble.Stats{
				Lines: 5,
				Words: 5,
				Bytes: 24,
			},
		},
		"only spaces": {
			input: "       ",
			want: gobble.Stats{
				Lines: 0,
				Words: 0,
				Bytes: 7,
			},
		},
		"multiple spaces between words": {
			input: "one two three  four five six",
			want: gobble.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 28,
			},
		},
		"utf8 spaces": {
			input: "one two three four five six",
			want: gobble.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 37,
			},
		},
		"unicode characters": {
			input: "Ђ ʩ",
			want: gobble.Stats{
				Lines: 0,
				Words: 2,
				Bytes: 5,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			got := gobble.Count(r)
			if got != tc.want {
				t.Logf("got: %v, want: %v", got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestPrint(t *testing.T) {
	type input struct {
		stats gobble.Stats
		label string
	}

	tests := map[string]struct {
		input input
		want  string
	}{
		"empty label": {
			input: input{
				stats: gobble.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: "",
			},
			want: "1 5 24\n",
		},
		"only words": {
			input: input{
				stats: gobble.Stats{
					Lines: 0,
					Words: 5,
					Bytes: 23,
				},
				label: "words.txt",
			},
			want: "0 5 23 words.txt\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := &bytes.Buffer{}
			tc.input.stats.Print(got, tc.input.label)

			if got.String() != tc.want {
				t.Logf("got %s, want: %s", got, tc.want)
				t.Fail()
			}
		})
	}
}
