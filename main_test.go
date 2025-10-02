package main

import "testing"

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
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CountWords([]byte(tc.input))
			if got != tc.want {
				t.Logf("got: %d, want: %d", got, tc.want)
				t.Fail()
			}
		})
	}
}
