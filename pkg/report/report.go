package report

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/d33trik/gobble/pkg/counter"
)

type Printer struct {
	writer io.Writer
	opts   Options
}

type Options struct {
	PrintHeader bool
	PrintLines  bool
	PrintWords  bool
	PrintBytes  bool
}

func NewPrinter(w io.Writer, opts Options) Printer {
	return Printer{
		writer: w,
		opts:   opts,
	}
}

func (p Printer) PrintHeader() {
	if !p.opts.PrintHeader {
		return
	}

	headers := []string{}

	if p.shouldPrintLines() {
		headers = append(headers, "lines")
	}

	if p.shouldPrintWords() {
		headers = append(headers, "words")
	}

	if p.shouldPrintBytes() {
		headers = append(headers, "bytes")
	}

	headersLine := strings.Join(headers, "\t")
	fmt.Fprintf(p.writer, "%s\t\n", headersLine)
}

func (p Printer) PrintStats(s counter.Stats, labels ...string) {
	counts := []string{}

	if p.shouldPrintLines() {
		counts = append(counts, strconv.Itoa(s.Lines))
	}

	if p.shouldPrintWords() {
		counts = append(counts, strconv.Itoa(s.Words))
	}

	if p.shouldPrintBytes() {
		counts = append(counts, strconv.Itoa(s.Bytes))
	}

	countsLine := strings.Join(counts, "\t")
	fmt.Fprintf(p.writer, "%s\t", countsLine)

	labelLine := strings.Join(labels, " ")
	if labelLine != "" {
		fmt.Fprintf(p.writer, " %s", labelLine)
	}

	fmt.Fprintf(p.writer, "\n")
}

func (p Printer) useDefault() bool {
	return !p.opts.PrintLines && !p.opts.PrintWords && !p.opts.PrintBytes
}

func (p Printer) shouldPrintLines() bool {
	return p.opts.PrintLines || p.useDefault()
}

func (p Printer) shouldPrintWords() bool {
	return p.opts.PrintWords || p.useDefault()
}

func (p Printer) shouldPrintBytes() bool {
	return p.opts.PrintBytes || p.useDefault()
}
