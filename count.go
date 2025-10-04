package main

import (
	"bufio"
	"io"
	"os"
)

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return CountWords(file), nil
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
