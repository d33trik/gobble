package main

import (
	"bufio"
	"io"
)

func Count(rs io.ReadSeeker) (lines, words, bytes int) {
	const offsetStart = 0

	lines = CountLines(rs)
	rs.Seek(offsetStart, io.SeekStart)

	words = CountWords(rs)
	rs.Seek(offsetStart, io.SeekStart)

	bytes = CountBytes(rs)

	return lines, words, bytes
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
