package extractor

import (
	"local/tedi/src/finder"
)

type extractor struct {
	folder  string
	ignore  []string
	pattern []string
	output  string
}

func (e extractor) findFiles() []string {
	ff := finder.Files{
		Path:   e.folder,
		Ignore: e.ignore,
	}
	return ff.Find()
}

func (e extractor) extract(file string) []string {
	var texts []string
	// Todo
	return texts
}

func (e extractor) ExtractAll() map[string][]string {
	extractedTexts := make(map[string][]string)
	files := e.findFiles()

	for _, file := range files {
		texts := e.extract(file)
		if texts == nil {
			continue
		}
		extractedTexts[file] = texts
	}
	return extractedTexts
}

func New(folder string, ignore []string, pattern []string, output string) extractor {
	e := extractor{
		folder:  folder,
		ignore:  ignore,
		pattern: pattern,
		output:  output,
	}
	return e
}
