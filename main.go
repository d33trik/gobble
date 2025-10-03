package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	data, err := os.ReadFile("./testdata/words.txt")
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	numberOfWords := CountWords(data)

	fmt.Println(numberOfWords)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
