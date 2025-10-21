package gobble

type DisplayOptions struct {
	Header bool
	Lines  bool
	Words  bool
	Bytes  bool
}

func (d DisplayOptions) printAll() bool {
	return !d.Lines && !d.Words && !d.Bytes
}

func (d DisplayOptions) printLines() bool {
	return d.Lines || d.printAll()
}

func (d DisplayOptions) printWords() bool {
	return d.Words || d.printAll()
}

func (d DisplayOptions) printBytes() bool {
	return d.Bytes || d.printAll()
}
