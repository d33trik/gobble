package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./testdata/words.txt")
	if err != nil {
		panic(err)
	}

	numberOfWords := CountWords(data)

	fmt.Println(numberOfWords)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
