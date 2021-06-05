package finder

type Finder interface {
	Find() []string
}

type files struct {
	path    string
	ignore  []string
	results []string
}

func Files(path string, ignore []string) files {
	return files{
		path:   path,
		ignore: ignore,
	}
}

type texts struct {
	text    string
	pattern []string
	results []string
}

func Texts(text string, pattern []string) texts {
	return texts{
		text:    text,
		pattern: pattern,
	}
}
