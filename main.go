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

	fmt.Println("data:", string(data))
}
