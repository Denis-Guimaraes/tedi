package replacer

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type replacer struct {
	path      []string
	ignore    []string
	delimiter string
	csv       string
}

func New(path []string, ignore []string, delimiter string, csv string) *replacer {
	e := &replacer{
		path:      path,
		ignore:    ignore,
		delimiter: delimiter,
		csv:       csv,
	}
	return e
}

func (r *replacer) ReplaceAll() {
	lines := r.readCsv()
	for _, line := range lines {
		r.replace(line)
	}
}

func (r *replacer) readCsv() [][]string {
	file := r.openCsv()
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error read csv: %v\n", err)
		os.Exit(1)
	}

	file.Close()
	return lines[1:]
}

func (r *replacer) openCsv() *os.File {
	file, err := os.Open(r.csv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error open csv: %v\n", err)
		os.Exit(1)
	}
	return file
}

func (r *replacer) replace(line []string) {
	for _, path := range r.path {
		fp := filepath.Join(path, line[0])
		content := r.readFile(fp)
		if content == nil {
			continue
		}

		oldText := fmt.Sprintf("%s%s%s", r.delimiter, line[1], r.delimiter)
		newContent := strings.ReplaceAll(string(content), oldText, line[2])
		r.WriteFile(fp, newContent)
	}
}

func (r *replacer) readFile(path string) []byte {
	text, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error read file: %v\n", err)
		return nil
	}
	return text
}

func (r *replacer) WriteFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "content changed in %s\n", path)
}
