package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatalln("error: no file provided")
	}

	total := 0
	hadError := false
	filenames := os.Args[1:]

	for _, filename := range filenames {
		numberOfWords, err := CountWordsInFile(filename)
		if err != nil {
			hadError = true
			fmt.Fprintln(os.Stderr, "gobble:", err)
			continue
		}

		total = total + numberOfWords

		fmt.Println(numberOfWords, filename)
	}

	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

	if hadError {
		os.Exit(1)
	}
}

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	return CountWords(file), nil
}

func CountWords(r io.Reader) (numberOfWords int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		numberOfWords++
	}

	return numberOfWords
}
