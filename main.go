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
	filenames := os.Args[1:]

	for _, filename := range filenames {
		numberOfWords := CountWordsInFile(filename)
		total = total + numberOfWords

		fmt.Println(numberOfWords, filename)
	}

	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}
}

func CountWordsInFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	return CountWords(file)
}

func CountWords(r io.Reader) (numberOfWords int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		numberOfWords++
	}

	return numberOfWords
}
