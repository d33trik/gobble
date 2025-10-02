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
	if len(data) == 0 {
		return 0
	}

	for _, byte := range data {
		if byte == ' ' {
			numberOfWords++
		}
	}

	numberOfWords++

	return numberOfWords
}
