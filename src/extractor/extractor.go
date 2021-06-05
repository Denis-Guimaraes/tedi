package extractor

import (
	"fmt"
	"io/ioutil"
	"local/tedi/src/finder"
	"os"
)

type extractor struct {
	folder  string
	ignore  []string
	pattern []string
	output  string
}

func (e extractor) findFiles() []string {
	ff := finder.Files(e.folder, e.ignore)
	return ff.Find()
}

func (e extractor) extract(file string) []string {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	tf := finder.Texts(string(text), e.pattern)
	return tf.Find()
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
