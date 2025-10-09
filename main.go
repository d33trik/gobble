package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

type DisplayOptions struct {
	PrintLines bool
	PrintWords bool
	PrintBytes bool
}

func (d DisplayOptions) UseDefault() bool {
	return !d.PrintLines && !d.PrintWords && !d.PrintBytes
}

func (d DisplayOptions) ShouldPrintLines() bool {
	return d.PrintLines || d.UseDefault()
}

func (d DisplayOptions) ShouldPrintWords() bool {
	return d.PrintWords || d.UseDefault()
}

func (d DisplayOptions) ShouldPrintBytes() bool {
	return d.PrintBytes || d.UseDefault()
}

func main() {
	opts := DisplayOptions{}
	log.SetFlags(0)

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	flag.BoolVar(&opts.PrintLines, "l", false, "Print the number of lines")
	flag.BoolVar(&opts.PrintWords, "w", false, "Print the number of words")
	flag.BoolVar(&opts.PrintBytes, "b", false, "Print the number of bytes")
	flag.Parse()

	totals := Stats{}
	hadError := false
	filenames := flag.Args()

	if len(filenames) == 0 {
		stats := Count(os.Stdin)
		stats.Print(tw, opts)
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

			stats.Print(tw, opts, filename)
		}()
	}

	if len(filenames) > 1 {
		totals.Print(tw, opts, "total")
	}

	tw.Flush()

	if hadError {
		os.Exit(1)
	}
}
