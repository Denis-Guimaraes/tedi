package finder

import (
	"fmt"
	"os"
	"path/filepath"
)

func (f Files) Find() []string {
	err := filepath.Walk(f.Path, f.check)
	if err != nil {
		fmt.Println(err)
	}
	return f.results
}

func (f Files) check(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if f.isIgnore(path) {
		if info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	}
	if !info.IsDir() {
		fmt.Println("file found", path)
		f.results = append(f.results, path)
	}
	return nil
}

func (f Files) isIgnore(path string) bool {
	isIgnore := false
	for _, pattern := range f.Ignore {
		match, err := filepath.Match(pattern, path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if match {
			isIgnore = true
			break
		}
	}
	return isIgnore
}
