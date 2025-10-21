package gobble

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Printer struct {
	writer io.Writer
	opts   DisplayOptions
}

func NewPrinter(w io.Writer, opts DisplayOptions) Printer {
	return Printer{
		writer: w,
		opts:   opts,
	}
}

func (p Printer) PrintHeader() {
	if !p.opts.Header {
		return
	}

	headers := []string{}

	if p.opts.printLines() {
		headers = append(headers, "lines")
	}

	if p.opts.printWords() {
		headers = append(headers, "words")
	}

	if p.opts.printBytes() {
		headers = append(headers, "bytes")
	}

	headersLine := strings.Join(headers, "\t")
	fmt.Fprintf(p.writer, "%s\t\n", headersLine)
}

func (p Printer) PrintStats(s Stats, labels ...string) {
	counts := []string{}

	if p.opts.printLines() {
		counts = append(counts, strconv.Itoa(s.Lines))
	}

	if p.opts.printWords() {
		counts = append(counts, strconv.Itoa(s.Words))
	}

	if p.opts.printBytes() {
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
