package gobble

import (
	"bufio"
	"io"
	"os"
	"sync"
	"unicode"
)

type FileStats struct {
	Stats    Stats
	Filename string
}

type Stats struct {
	Lines int
	Words int
	Bytes int
}

func (s *Stats) Add(other Stats) {
	s.Lines += other.Lines
	s.Words += other.Words
	s.Bytes += other.Bytes
}

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

func CountFiles(filenames []string) (<-chan FileStats, <-chan error) {
	statsCh := make(chan FileStats)
	errCh := make(chan error)

	wg := sync.WaitGroup{}
	wg.Add(len(filenames))

	for _, filename := range filenames {
		go func() {
			defer wg.Done()

			file, err := os.Open(filename)
			if err != nil {
				errCh <- err
				return
			}
			defer file.Close()

			statsCh <- FileStats{
				Stats:    Count(file),
				Filename: filename,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(statsCh)
		close(errCh)
	}()

	return statsCh, errCh
}
