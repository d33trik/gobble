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
		"empty input": {
			input: "",
			want:  0,
		},
		"no new lines": {
			input: "one two three four five six",
			want:  0,
		},
		"one new line": {
			input: "one two three four five\n",
			want:  1,
		},
		"no new lines at end": {
			input: "one two three four five\nsix",
			want:  1,
		},
		"multiple lines with words": {
			input: "one\ntow\nthree\nfour\nfive\n",
			want:  5,
		},
		"only lines": {
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
		"empty input": {
			input: "",
			want:  0,
		},
		"five words": {
			input: "one two three four five",
			want:  5,
		},
		"single space": {
			input: " ",
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

func TestCountBytes(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"empty input": {
			input: "",
			want:  0,
		},
		"five words": {
			input: "one two three four five",
			want:  23,
		},
		"only spaces": {
			input: "       ",
			want:  7,
		},
		"multiple lines with words": {
			input: "one\ntwo\nthree\nfour\n",
			want:  19,
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
		input     string
		wantLines int
		wantWords int
		wantBytes int
	}{
		"empty input": {
			input:     "",
			wantLines: 0,
			wantWords: 0,
			wantBytes: 0,
		},
		"five words": {
			input:     "one two three four five",
			wantLines: 0,
			wantWords: 5,
			wantBytes: 23,
		},
		"multiple lines with words": {
			input:     "one\ntwo\nthree\nfour\nfive\n",
			wantLines: 5,
			wantWords: 5,
			wantBytes: 24,
		},
		"only lines": {
			input:     "\n\n\n\n\n",
			wantLines: 5,
			wantWords: 0,
			wantBytes: 5,
		},
		"only spaces": {
			input:     "       ",
			wantLines: 0,
			wantWords: 0,
			wantBytes: 7,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			gotLines, gotWords, gotBytes := gobble.Count(r)
			if gotLines != tc.wantLines || gotWords != tc.wantWords || gotBytes != tc.wantBytes {
				t.Logf("got: (%d, %d, %d), want: (%d, %d, %d)", gotLines, gotWords, gotBytes, tc.wantLines, tc.wantWords, tc.wantBytes)
				t.Fail()
			}
		})
	}
}
