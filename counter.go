package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type Stats struct {
	Lines int
	Words int
	Bytes int
}

func (s *Stats) Print(w io.Writer, opts DisplayOptions, labels ...string) {
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

func (s *Stats) Add(other Stats) {
	s.Lines += other.Lines
	s.Words += other.Words
	s.Bytes += other.Bytes
}

func Count(r io.Reader) (stats Stats) {
	isInsideWord := false
	reader := bufio.NewReader(r)

	for {
		rune, size, err := reader.ReadRune()
		if err != nil {
			break
		}

		if rune == '\n' {
			stats.Lines++
		}

		isSpace := unicode.IsSpace(rune)

		if !isSpace && !isInsideWord {
			stats.Words++
		}

		isInsideWord = !isSpace

		stats.Bytes += size
	}

	return stats
}
