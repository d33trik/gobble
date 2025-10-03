package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	file, err := os.Open("./testdata/utf8.txt")
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	numberOfWords := CountWordsInFile(file)

	fmt.Println(numberOfWords)
}

func CountWordsInFile(file *os.File) (numberOfWords int) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		numberOfWords++
	}

	return numberOfWords
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
