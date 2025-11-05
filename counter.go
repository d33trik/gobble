package gobble

import (
	"bufio"
	"io"
	"os"
	"sync"
	"unicode"
)

func Count(r io.Reader) (stats Stats) {
	isInsideWord := false
	reader := bufio.NewReader(r)

	for {
		rune, size, err := reader.ReadRune()
		if err != nil {
			break
		}

		if rune == '\n' {
			stats.Lines++
		}

		isSpace := unicode.IsSpace(rune)

		if !isSpace && !isInsideWord {
			stats.Words++
		}

		isInsideWord = !isSpace

		stats.Bytes += size
	}

	return stats
}

func CountFiles(filenames []string) <-chan Stats {
	ch := make(chan Stats)

	wg := sync.WaitGroup{}
	wg.Add(len(filenames))

	for _, filename := range filenames {
		go func() {
			defer wg.Done()

			stats := Stats{}

			file, err := os.Open(filename)
			if err != nil {
				stats.Err = err
				ch <- stats
				return
			}
			defer file.Close()

			stats = Count(file)
			stats.Filename = filename

			ch <- stats
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
