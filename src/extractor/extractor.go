package extractor

import (
	"fmt"
	"local/tedi/src/finder"
	"local/tedi/src/logger"
	"os"
)

type extractor struct {
	path      string
	extension []string
	ignore    []string
	delimiter string
}

func New(path string, extension []string, ignore []string, delimiter string) *extractor {
	e := &extractor{
		path:      path,
		extension: extension,
		ignore:    ignore,
		delimiter: delimiter,
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
	ff := finder.File(e.path, e.extension, e.ignore)
	files = append(files, ff.Find()...)
	return files
}

func (e *extractor) extract(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		logger.Error(fmt.Sprintf("error: %v", err))
		os.Exit(1)
	}
	tf := finder.Text(string(content), e.delimiter)
	return tf.Find()
}
