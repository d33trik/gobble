package counter_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/d33trik/gobble/pkg/counter"
	"github.com/d33trik/gobble/pkg/display"
)

func TestPrint(t *testing.T) {
	type input struct {
		stats counter.Stats
		label []string
	}

	tests := map[string]struct {
		input input
		opts  display.Options
		want  string
	}{
		"empty label": {
			input: input{
				stats: counter.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{},
			},
			opts: display.Options{
				PrintLines: false,
				PrintWords: false,
				PrintBytes: false,
			},
			want: "1\t5\t24\t\n",
		},
		"print default": {
			input: input{
				stats: counter.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: display.Options{
				PrintLines: false,
				PrintWords: false,
				PrintBytes: false,
			},
			want: "1\t5\t24\t file.txt\n",
		},
		"print all": {
			input: input{
				stats: counter.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: display.Options{
				PrintLines: true,
				PrintWords: true,
				PrintBytes: true,
			},
			want: "1\t5\t24\t file.txt\n",
		},
		"print only lines": {
			input: input{
				stats: counter.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: display.Options{
				PrintLines: true,
				PrintWords: false,
				PrintBytes: false,
			},
			want: "1\t file.txt\n",
		},
		"print only words": {
			input: input{
				stats: counter.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: display.Options{
				PrintLines: false,
				PrintWords: true,
				PrintBytes: false,
			},
			want: "5\t file.txt\n",
		},
		"print only bytes": {
			input: input{
				stats: counter.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: display.Options{
				PrintLines: false,
				PrintWords: false,
				PrintBytes: true,
			},
			want: "24\t file.txt\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := &bytes.Buffer{}
			tc.input.stats.Print(got, tc.opts, tc.input.label...)

			if got.String() != tc.want {
				t.Logf("got %s, want: %s", got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type input struct {
		stats counter.Stats
		other counter.Stats
	}

	tests := map[string]struct {
		input input
		want  counter.Stats
	}{
		"simple add": {
			input: input{
				stats: counter.Stats{
					Lines: 5,
					Words: 5,
					Bytes: 24,
				},
				other: counter.Stats{
					Lines: 1,
					Words: 1,
					Bytes: 4,
				},
			},
			want: counter.Stats{
				Lines: 6,
				Words: 6,
				Bytes: 28,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.stats
			got.Add(tc.input.other)

			if got != tc.want {
				t.Logf("got: %v, want: %v", got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := map[string]struct {
		input string
		want  counter.Stats
	}{
		"empty input": {
			input: "",
			want: counter.Stats{
				Lines: 0,
				Words: 0,
				Bytes: 0,
			},
		},
		"only words": {
			input: "one two three four five six",
			want: counter.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 27,
			},
		},
		"only new lines": {
			input: "\n\n\n\n\n",
			want: counter.Stats{
				Lines: 5,
				Words: 0,
				Bytes: 5,
			},
		},
		"one new line": {
			input: "one two three\nfour five six",
			want: counter.Stats{
				Lines: 1,
				Words: 6,
				Bytes: 27,
			},
		},
		"start one new line": {
			input: "\none two three four five six",
			want: counter.Stats{
				Lines: 1,
				Words: 6,
				Bytes: 28,
			},
		},
		"end one new line": {
			input: "one two three four five six\n",
			want: counter.Stats{
				Lines: 1,
				Words: 6,
				Bytes: 28,
			},
		},
		"multiple new lines": {
			input: "one\ntwo\nthree\nfour\nfive\nsix\n",
			want: counter.Stats{
				Lines: 6,
				Words: 6,
				Bytes: 28,
			},
		},
		"only spaces": {
			input: "       ",
			want: counter.Stats{
				Lines: 0,
				Words: 0,
				Bytes: 7,
			},
		},
		"start multiple spaces": {
			input: "  one two three four five six",
			want: counter.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 29,
			},
		},
		"end multiple spaces": {
			input: "one two three four five six  ",
			want: counter.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 29,
			},
		},
		"multiple spaces between words": {
			input: "one two three  four five six",
			want: counter.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 28,
			},
		},
		"utf8 spaces": {
			input: "one two three four five six",
			want: counter.Stats{
				Lines: 0,
				Words: 6,
				Bytes: 37,
			},
		},
		"unicode characters": {
			input: "Ђ ʩ",
			want: counter.Stats{
				Lines: 0,
				Words: 2,
				Bytes: 5,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			got := counter.Count(r)
			if got != tc.want {
				t.Logf("got: %v, want: %v", got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestPrintHeader(t *testing.T) {
	tests := map[string]struct {
		opts display.Options
		want string
	}{
		"print all headers": {
			opts: display.Options{
				PrintLines: true,
				PrintWords: true,
				PrintBytes: true,
			},
			want: "lines\twords\tbytes\t\n",
		},
		"print only lines header": {
			opts: display.Options{
				PrintLines: true,
				PrintWords: false,
				PrintBytes: false,
			},
			want: "lines\t\n",
		},
		"print only words header": {
			opts: display.Options{
				PrintLines: false,
				PrintWords: true,
				PrintBytes: false,
			},
			want: "words\t\n",
		},
		"print only bytes header": {
			opts: display.Options{
				PrintLines: false,
				PrintWords: false,
				PrintBytes: true,
			},
			want: "bytes\t\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := &bytes.Buffer{}
			counter.PrintHeader(got, tc.opts)

			if got.String() != tc.want {
				t.Logf("got %s, want: %s", got, tc.want)
				t.Fail()
			}
		})
	}
}
