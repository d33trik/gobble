package gobble

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
