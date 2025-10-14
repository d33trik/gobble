package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/d33trik/gobble/pkg/counter"
	"github.com/d33trik/gobble/pkg/report"
)

func main() {
	log.SetFlags(0)
	opts := parseFlags()

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)
	printer := report.NewPrinter(tw, opts)
	printer.PrintHeader()

	hadError := false
	totals := counter.Stats{}
	filenames := flag.Args()

	if len(filenames) == 0 {
		stats := counter.Count(os.Stdin)
		printer.PrintStats(stats)
	}

	ch, errCh := counter.CountFiles(filenames)

	for ch != nil || errCh != nil {
		select {
		case fileStats, open := <-ch:
			if !open {
				ch = nil
				continue
			}

			totals.Add(fileStats.Stats)
			printer.PrintStats(fileStats.Stats, fileStats.Filename)
		case err, open := <-errCh:
			if !open {
				errCh = nil
				continue
			}

			hadError = true
			fmt.Fprintln(os.Stderr, "gobble:", err)
		}
	}

	if len(filenames) > 1 {
		printer.PrintStats(totals, "total")
	}

	tw.Flush()

	if hadError {
		os.Exit(1)
	}
}

func parseFlags() report.Options {
	opts := report.Options{}

	flag.BoolVar(&opts.PrintHeader, "h", false, "Print header")
	flag.BoolVar(&opts.PrintLines, "l", false, "Print the number of new lines")
	flag.BoolVar(&opts.PrintWords, "w", false, "Print the number of words")
	flag.BoolVar(&opts.PrintBytes, "b", false, "Print the number of bytes")
	flag.Parse()

	return opts
}
