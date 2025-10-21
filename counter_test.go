package gobble_test

import (
	"strings"
	"testing"

	"github.com/d33trik/gobble"
)

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
			input: "one two three four five six",
			want: gobble.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 27,
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
		"one new line": {
			input: "one two three\nfour five six",
			want: gobble.Stats{
				Lines: 1,
				Words: 6,
				Bytes: 27,
			},
		},
		"start one new line": {
			input: "\none two three four five six",
			want: gobble.Stats{
				Lines: 1,
				Words: 6,
				Bytes: 28,
			},
		},
		"end one new line": {
			input: "one two three four five six\n",
			want: gobble.Stats{
				Lines: 1,
				Words: 6,
				Bytes: 28,
			},
		},
		"multiple new lines": {
			input: "one\ntwo\nthree\nfour\nfive\nsix\n",
			want: gobble.Stats{
				Lines: 6,
				Words: 6,
				Bytes: 28,
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
		"start multiple spaces": {
			input: "  one two three four five six",
			want: gobble.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 29,
			},
		},
		"end multiple spaces": {
			input: "one two three four five six  ",
			want: gobble.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 29,
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
