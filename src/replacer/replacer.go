package replacer

import (
	"encoding/csv"
	"fmt"
	"local/tedi/src/logger"
	"os"
	"path/filepath"
	"strings"
)

type replacer struct {
	path      string
	delimiter string
	csv       string
}

func New(path string, delimiter string, csv string) *replacer {
	e := &replacer{
		path:      path,
		delimiter: delimiter,
		csv:       csv,
	}
	return e
}

func (r *replacer) ReplaceAll() {
	lines := r.readCsv()
	for index, line := range lines {
		if !r.isValideLine(line, index+2) {
			continue
		}
		r.replace(line)
	}
}

func (r *replacer) readCsv() [][]string {
	file := r.openCsv()
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		logger.Error(fmt.Sprintf("error read csv %s: %v", r.csv, err))
		os.Exit(1)
	}

	file.Close()
	return lines[1:]
}

func (r *replacer) openCsv() *os.File {
	file, err := os.Open(r.csv)
	if err != nil {
		logger.Error(fmt.Sprintf("error open csv %s: %v", r.csv, err))
		os.Exit(1)
	}
	return file
}

func (r *replacer) isValideLine(line []string, lineNumber int) bool {
	if len(line[0]) == 0 {
		logger.Error(fmt.Sprintf("line %d: file column empty", lineNumber))
		return false
	}
	if len(line[1]) == 0 {
		logger.Error(fmt.Sprintf("line %d: text column empty", lineNumber))
		return false
	}
	if len(line[2]) == 0 {
		logger.Error(fmt.Sprintf("line %d: replace column empty", lineNumber))
		return false
	}
	return true
}

func (r *replacer) replace(line []string) {
	fp := filepath.Join(r.path, line[0])
	content := r.readFile(fp)
	if content == nil {
		return
	}

	oldText := fmt.Sprintf("%s%s%s", r.delimiter, line[1], r.delimiter)
	if !strings.Contains(string(content), oldText) {
		logger.Error(fmt.Sprintf("note found text '%s' in file %s", oldText, fp))
		return
	}

	newContent := strings.ReplaceAll(string(content), oldText, line[2])
	success := r.WriteFile(fp, newContent)
	if success {
		logger.Info(fmt.Sprintf("text change from '%s' to '%s' in file %s", oldText, line[2], fp))
	}
}

func (r *replacer) readFile(path string) []byte {
	text, err := os.ReadFile(path)
	if err != nil {
		logger.Error(fmt.Sprintf("error read file %s: %v", path, err))
		return nil
	}
	return text
}

func (r *replacer) WriteFile(path string, content string) bool {
	err := os.WriteFile(path, []byte(content), 0)
	if err != nil {
		logger.Error(fmt.Sprintf("error write file %s: %v", path, err))
		return false
	}
	return true
}
