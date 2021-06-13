package finder

import (
	"fmt"
	"os"
	"regexp"
)

func (t *text) Find() []string {
	for _, pattern := range t.pattern {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		matches := regex.FindAllStringSubmatch(string(t.content), -1)
		t.addResults(matches)
	}
	return t.results
}

func (t *text) addResults(matches [][]string) {
	for _, match := range matches {
		t.results = append(t.results, match[1])
	}
}
