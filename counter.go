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
	fields := []string{}

	if opts.ShouldPrintLines() {
		fields = append(fields, strconv.Itoa(s.Lines))
	}

	if opts.ShouldPrintWords() {
		fields = append(fields, strconv.Itoa(s.Words))
	}

	if opts.ShouldPrintBytes() {
		fields = append(fields, strconv.Itoa(s.Bytes))
	}

	fields = append(fields, labels...)

	line := strings.Join(fields, " ")

	fmt.Fprintln(w, line)
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
		r, size, err := reader.ReadRune()
		if err != nil {
			break
		}

		if r == '\n' {
			stats.Lines++
		}

		isSpace := unicode.IsSpace(r)

		if !isSpace && !isInsideWord {
			stats.Words++
		}

		isInsideWord = !isSpace

		stats.Bytes += size
	}

	return stats
}
