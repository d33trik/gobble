package gobble

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Stats struct {
	Lines    int
	Words    int
	Bytes    int
	Filename string
	Err      error
}

func (s *Stats) Add(other Stats) {
	s.Lines += other.Lines
	s.Words += other.Words
	s.Bytes += other.Bytes
}

func (s *Stats) Print(w io.Writer, opts DisplayOptions, labels ...string) {
	counts := []string{}

	if opts.printLines() {
		counts = append(counts, strconv.Itoa(s.Lines))
	}

	if opts.printWords() {
		counts = append(counts, strconv.Itoa(s.Words))
	}

	if opts.printBytes() {
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

func PrintHeader(w io.Writer, opts DisplayOptions) {
	if !opts.Header {
		return
	}

	headers := []string{}

	if opts.printLines() {
		headers = append(headers, "lines")
	}

	if opts.printWords() {
		headers = append(headers, "words")
	}

	if opts.printBytes() {
		headers = append(headers, "bytes")
	}

	headersLine := strings.Join(headers, "\t")
	fmt.Fprintf(w, "%s\t\n", headersLine)
}
