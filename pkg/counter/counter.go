package counter

import (
	"bufio"
	"io"
	"unicode"
)

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
