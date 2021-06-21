package extractor

import (
	"fmt"
	"io/ioutil"
	"local/tedi/src/finder"
	"os"
)

type extractor struct {
	folder  []string
	ignore  []string
	pattern []string
}

func New(folder []string, ignore []string, pattern []string) *extractor {
	e := &extractor{
		folder:  folder,
		ignore:  ignore,
		pattern: pattern,
	}
	return e
}

func (e *extractor) ExtractAll() map[string][]string {
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

func (e *extractor) findFiles() []string {
	var files []string
	for _,folder := range e.folder {
		ff := finder.File(folder, e.ignore)
		files = append(files, ff.Find()...)
	}
	return files
}

func (e *extractor) extract(file string) []string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	tf := finder.Text(string(content), e.pattern)
	return tf.Find()
}
