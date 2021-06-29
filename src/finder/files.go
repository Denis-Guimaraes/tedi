package finder

import (
	"fmt"
	"os"
	"path/filepath"
)

func (f *file) Find() []string {
	err := filepath.Walk(f.path, f.check)
	if err != nil {
		f.error(err)
	}
	return f.results
}

func (f *file) check(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if f.isIgnore(path) {
		if info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	}
	if !info.IsDir() && f.matchExtension(path) {
		f.results = append(f.results, path)
	}
	return nil
}

func (f *file) matchExtension(path string) bool {
	match := false
	fileExtension := filepath.Ext(path)
	for _, extension := range f.extension {
		if fileExtension == extension {
			match = true
			break
		}
	}
	return match
}

func (f *file) isIgnore(path string) bool {
	isIgnore := false
	for _, pattern := range f.ignore {
		match, err := filepath.Match(pattern, path)
		if err != nil {
			f.error(err)
		}
		if match {
			isIgnore = true
			break
		}
	}
	return isIgnore
}

func (f *file) error(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
