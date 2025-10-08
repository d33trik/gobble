package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
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

func Count(rs io.ReadSeeker) Stats {
	const offsetStart = 0

	lines := CountLines(rs)
	rs.Seek(offsetStart, io.SeekStart)

	words := CountWords(rs)
	rs.Seek(offsetStart, io.SeekStart)

	bytes := CountBytes(rs)

	return Stats{
		Lines: lines,
		Words: words,
		Bytes: bytes,
	}
}

func CountLines(r io.Reader) (numberOfLines int) {
	reader := bufio.NewReader(r)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if r == '\n' {
			numberOfLines++
		}
	}

	return numberOfLines
}

func CountWords(r io.Reader) (numberOfWords int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		numberOfWords++
	}

	return numberOfWords
}

func CountBytes(r io.Reader) (numberOfBytes int) {
	reader := bufio.NewReader(r)

	for {
		_, err := reader.ReadByte()
		if err != nil {
			break
		}

		numberOfBytes++
	}

	return numberOfBytes
}
