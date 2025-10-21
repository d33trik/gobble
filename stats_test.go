package gobble_test

import (
	"bytes"
	"testing"

	"github.com/d33trik/gobble"
)

func TestAdd(t *testing.T) {
	type input struct {
		stats gobble.Stats
		other gobble.Stats
	}

	tests := map[string]struct {
		input input
		want  gobble.Stats
	}{
		"simple add": {
			input: input{
				stats: gobble.Stats{
					Lines: 5,
					Words: 5,
					Bytes: 24,
				},
				other: gobble.Stats{
					Lines: 1,
					Words: 1,
					Bytes: 4,
				},
			},
			want: gobble.Stats{
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

func TestPrintStats(t *testing.T) {
	type input struct {
		stats gobble.Stats
		label []string
	}

	tests := map[string]struct {
		input input
		opts  gobble.DisplayOptions
		want  string
	}{
		"empty label": {
			input: input{
				stats: gobble.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{},
			},
			opts: gobble.DisplayOptions{
				Lines: false,
				Words: false,
				Bytes: false,
			},
			want: "1\t5\t24\t\n",
		},
		"print default": {
			input: input{
				stats: gobble.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: gobble.DisplayOptions{
				Lines: false,
				Words: false,
				Bytes: false,
			},
			want: "1\t5\t24\t file.txt\n",
		},
		"print all": {
			input: input{
				stats: gobble.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: gobble.DisplayOptions{
				Lines: true,
				Words: true,
				Bytes: true,
			},
			want: "1\t5\t24\t file.txt\n",
		},
		"print only lines": {
			input: input{
				stats: gobble.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: gobble.DisplayOptions{
				Lines: true,
				Words: false,
				Bytes: false,
			},
			want: "1\t file.txt\n",
		},
		"print only words": {
			input: input{
				stats: gobble.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: gobble.DisplayOptions{
				Lines: false,
				Words: true,
				Bytes: false,
			},
			want: "5\t file.txt\n",
		},
		"print only bytes": {
			input: input{
				stats: gobble.Stats{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				label: []string{"file.txt"},
			},
			opts: gobble.DisplayOptions{
				Lines: false,
				Words: false,
				Bytes: true,
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

func TestPrintHeader(t *testing.T) {
	tests := map[string]struct {
		opts gobble.DisplayOptions
		want string
	}{
		"print all headers": {
			opts: gobble.DisplayOptions{
				Header: true,
				Lines:  true,
				Words:  true,
				Bytes:  true,
			},
			want: "lines\twords\tbytes\t\n",
		},
		"print only lines header": {
			opts: gobble.DisplayOptions{
				Header: true,
				Lines:  true,
				Words:  false,
				Bytes:  false,
			},
			want: "lines\t\n",
		},
		"print only words header": {
			opts: gobble.DisplayOptions{
				Header: true,
				Lines:  false,
				Words:  true,
				Bytes:  false,
			},
			want: "words\t\n",
		},
		"print only bytes header": {
			opts: gobble.DisplayOptions{
				Header: true,
				Lines:  false,
				Words:  false,
				Bytes:  true,
			},
			want: "bytes\t\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := &bytes.Buffer{}
			gobble.PrintHeader(got, tc.opts)

			if got.String() != tc.want {
				t.Logf("got %s, want: %s", got, tc.want)
				t.Fail()
			}
		})
	}
}
