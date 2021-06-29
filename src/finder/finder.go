package finder

type Finder interface {
	Find() []string
}

type file struct {
	path      string
	extension []string
	ignore    []string
	results   []string
}

func File(path string, extension []string, ignore []string) *file {
	return &file{
		path:      path,
		extension: extension,
		ignore:    ignore,
	}
}

type text struct {
	content string
	pattern []string
	results []string
}

func Text(content string, pattern []string) *text {
	return &text{
		content: content,
		pattern: pattern,
	}
}
