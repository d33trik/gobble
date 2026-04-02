package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"codeberg.org/d33trik/gobble"
)

func main() {
	log.SetFlags(0)
	opts := parseFlags()

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)
	gobble.PrintHeader(tw, opts)

	hadError := false
	totals := gobble.Stats{}
	filenames := flag.Args()

	if len(filenames) == 0 {
		stats := gobble.Count(os.Stdin)
		stats.Print(tw, opts)
	}

	statsCh := gobble.CountFiles(filenames)

	for stats := range statsCh {
		if stats.Err != nil {
			hadError = true
			fmt.Fprintln(os.Stderr, "gobble:", stats.Err)
			continue
		}

		totals.Add(stats)
		stats.Print(tw, opts, stats.Filename)
	}

	if len(filenames) > 1 {
		totals.Print(tw, opts, "total")
	}

	tw.Flush()

	if hadError {
		os.Exit(1)
	}
}

func parseFlags() gobble.DisplayOptions {
	opts := gobble.DisplayOptions{}

	flag.BoolVar(&opts.Header, "h", false, "Print header")
	flag.BoolVar(&opts.Lines, "l", false, "Print the number of new lines")
	flag.BoolVar(&opts.Words, "w", false, "Print the number of words")
	flag.BoolVar(&opts.Bytes, "b", false, "Print the number of bytes")
	flag.Parse()

	return opts
}
