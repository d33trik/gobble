package report_test

import (
	"bytes"
	"testing"

	"github.com/d33trik/gobble/pkg/counter"
	"github.com/d33trik/gobble/pkg/report"
)

func TestPrintHeader(t *testing.T) {
	tests := map[string]struct {
		opts report.Options
		want string
	}{
		"print all headers": {
			opts: report.Options{
				PrintHeader: true,
				PrintLines:  true,
				PrintWords:  true,
				PrintBytes:  true,
			},
			want: "lines\twords\tbytes\t\n",
		},
		"print only lines header": {
			opts: report.Options{
				PrintHeader: true,
				PrintLines:  true,
				PrintWords:  false,
				PrintBytes:  false,
			},
			want: "lines\t\n",
		},
		"print only words header": {
			opts: report.Options{
				PrintHeader: true,
				PrintLines:  false,
				PrintWords:  true,
				PrintBytes:  false,
			},
			want: "words\t\n",
		},
		"print only bytes header": {
			opts: report.Options{
				PrintHeader: true,
				PrintLines:  false,
				PrintWords:  false,
				PrintBytes:  true,
			},
			want: "bytes\t\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := &bytes.Buffer{}
			report := report.NewPrinter(got, tc.opts)
			report.PrintHeader()

			if got.String() != tc.want {
				t.Logf("got %s, want: %s", got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestPrintStats(t *testing.T) {
	type input struct {
		stats counter.Stats
		label []string
	}

	tests := map[string]struct {
		input input
		opts  report.Options
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
			opts: report.Options{
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
			opts: report.Options{
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
			opts: report.Options{
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
			opts: report.Options{
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
			opts: report.Options{
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
			opts: report.Options{
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
			report := report.NewPrinter(got, tc.opts)
			report.PrintStats(tc.input.stats, tc.input.label...)

			if got.String() != tc.want {
				t.Logf("got %s, want: %s", got, tc.want)
				t.Fail()
			}
		})
	}
}
