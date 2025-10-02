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

	numberOfWords := CountWords(data)

	fmt.Println(numberOfWords)
}

func CountWords(data []byte) (numberOfWords int) {
	wasSpace := true

	for _, byte := range data {
		isSpace := (byte == ' ' || byte == '\n')

		if wasSpace && !isSpace {
			numberOfWords++
		}

		wasSpace = isSpace
	}

	return numberOfWords
}
