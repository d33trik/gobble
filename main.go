package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./testdata/words.txt")
	if err != nil {
		panic(err)
	}

	wordCount := countWords(data)

	fmt.Println(wordCount)
}

func countWords(data []byte) int {
	wordCount := 0

	for _, byte := range data {
		if byte == ' ' {
			wordCount++
		}
	}

	wordCount++

	return wordCount
}
