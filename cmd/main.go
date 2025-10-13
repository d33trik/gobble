package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/d33trik/gobble/pkg/counter"
	"github.com/d33trik/gobble/pkg/report"
)

func main() {
	log.SetFlags(0)
	opts := report.Options{}

	flag.BoolVar(&opts.PrintHeader, "h", false, "Print header")
	flag.BoolVar(&opts.PrintLines, "l", false, "Print the number of new lines")
	flag.BoolVar(&opts.PrintWords, "w", false, "Print the number of words")
	flag.BoolVar(&opts.PrintBytes, "b", false, "Print the number of bytes")
	flag.Parse()

	hadError := false
	totals := counter.Stats{}
	filenames := flag.Args()
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)
	printer := report.NewPrinter(tw, opts)

	printer.PrintHeader()

	if len(filenames) == 0 {
		stats := counter.Count(os.Stdin)
		printer.PrintStats(stats)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(filenames))

	ch := make(chan counter.FileStats)

	for _, filename := range filenames {
		go func() {
			defer wg.Done()

			file, err := os.Open(filename)
			if err != nil {
				hadError = true
				fmt.Fprintln(os.Stderr, "gobble:", err)
				return
			}
			defer file.Close()

			ch <- counter.FileStats{
				Stats:    counter.Count(file),
				Filename: filename,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for fileStats := range ch {
		totals.Add(fileStats.Stats)
		printer.PrintStats(fileStats.Stats, fileStats.Filename)
	}

	if len(filenames) > 1 {
		printer.PrintStats(totals, "total")
	}

	tw.Flush()

	if hadError {
		os.Exit(1)
	}
}
