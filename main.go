package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	total := 0
	hadError := false
	filenames := os.Args[1:]

	if len(filenames) == 0 {
		lines, words, bytes := Count(os.Stdin)
		fmt.Println(lines, words, bytes)
	}

	for _, filename := range filenames {
		func() {
			file, err := os.Open(filename)
			if err != nil {
				hadError = true
				fmt.Fprintln(os.Stderr, "gobble:", err)
				return
			}
			defer file.Close()

			lines, words, bytes := Count(file)
			total = total + words

			fmt.Println(lines, words, bytes, filename)
		}()
	}

	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

	if hadError {
		os.Exit(1)
	}
}
