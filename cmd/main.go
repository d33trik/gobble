package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/d33trik/gobble/pkg/counter"
	"github.com/d33trik/gobble/pkg/display"
)

func main() {
	opts := display.Options{}
	log.SetFlags(0)

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	flag.BoolVar(&opts.PrintHeader, "h", false, "Print header")
	flag.BoolVar(&opts.PrintLines, "l", false, "Print the number of lines")
	flag.BoolVar(&opts.PrintWords, "w", false, "Print the number of words")
	flag.BoolVar(&opts.PrintBytes, "b", false, "Print the number of bytes")
	flag.Parse()

	totals := counter.Stats{}
	hadError := false
	filenames := flag.Args()

	if opts.PrintHeader {
		display.PrintHeader(tw, opts)
	}

	if len(filenames) == 0 {
		stats := counter.Count(os.Stdin)
		display.PrintStats(tw, opts, stats)
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

			stats := counter.Count(file)
			totals.Add(stats)

			display.PrintStats(tw, opts, stats, filename)
		}()
	}

	if len(filenames) > 1 {
		display.PrintStats(tw, opts, totals, "total")
	}

	tw.Flush()

	if hadError {
		os.Exit(1)
	}
}
