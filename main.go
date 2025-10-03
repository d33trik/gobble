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

	file, err := os.Open("./testdata/utf8.txt")
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	numberOfWords := CountWords(file)

	fmt.Println(numberOfWords)
}

func CountWords(r io.Reader) (numberOfWords int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		numberOfWords++
	}

	return numberOfWords
}
