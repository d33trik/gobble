package display

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/d33trik/gobble/pkg/counter"
)

type Options struct {
	PrintHeader bool
	PrintLines  bool
	PrintWords  bool
	PrintBytes  bool
}

func (o Options) UseDefault() bool {
	return !o.PrintLines && !o.PrintWords && !o.PrintBytes
}

func (o Options) ShouldPrintLines() bool {
	return o.PrintLines || o.UseDefault()
}

func (o Options) ShouldPrintWords() bool {
	return o.PrintWords || o.UseDefault()
}

func (o Options) ShouldPrintBytes() bool {
	return o.PrintBytes || o.UseDefault()
}

func PrintHeader(w io.Writer, opts Options) {
	headers := []string{}

	if opts.ShouldPrintLines() {
		headers = append(headers, "lines")
	}

	if opts.ShouldPrintWords() {
		headers = append(headers, "words")
	}

	if opts.ShouldPrintBytes() {
		headers = append(headers, "bytes")
	}

	headersLine := strings.Join(headers, "\t")
	fmt.Fprintf(w, "%s\t\n", headersLine)
}

func PrintStats(w io.Writer, opts Options, s counter.Stats, labels ...string) {
	counts := []string{}

	if opts.ShouldPrintLines() {
		counts = append(counts, strconv.Itoa(s.Lines))
	}

	if opts.ShouldPrintWords() {
		counts = append(counts, strconv.Itoa(s.Words))
	}

	if opts.ShouldPrintBytes() {
		counts = append(counts, strconv.Itoa(s.Bytes))
	}

	countsLine := strings.Join(counts, "\t")
	fmt.Fprintf(w, "%s\t", countsLine)

	labelLine := strings.Join(labels, " ")
	if labelLine != "" {
		fmt.Fprintf(w, " %s", labelLine)
	}

	fmt.Fprintf(w, "\n")
}
