package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	totals := Stats{}
	hadError := false
	filenames := os.Args[1:]

	if len(filenames) == 0 {
		stats := Count(os.Stdin)
		stats.Print(os.Stdout)
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

			stats := Count(file)
			totals.Add(stats)

			stats.Print(os.Stdout, filename)
		}()
	}

	if len(filenames) > 1 {
		totals.Print(os.Stdout, "total")
	}

	if hadError {
		os.Exit(1)
	}
}
