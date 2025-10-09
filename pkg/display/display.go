package display

type Options struct {
	PrintHeader bool
	PrintLines  bool
	PrintWords  bool
	PrintBytes  bool
}

func (o Options) UseDefault() bool {
	return !o.PrintLines && !o.PrintWords && !o.PrintBytes
}

func (o Options) ShouldPrintLines() bool {
	return o.PrintLines || o.UseDefault()
}

func (o Options) ShouldPrintWords() bool {
	return o.PrintWords || o.UseDefault()
}

func (o Options) ShouldPrintBytes() bool {
	return o.PrintBytes || o.UseDefault()
}
