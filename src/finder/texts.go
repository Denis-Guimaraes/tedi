package finder

import (
	"fmt"
	"os"
	"regexp"
)

func (t texts) Find() []string {
	for _, pattern := range t.pattern {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		matches := regex.FindAllStringSubmatch(string(t.text), -1)
		t.addResults(matches)
	}
	return t.results
}

func (t *texts) addResults(matches [][]string) {
	for _, match := range matches {
		t.results = append(t.results, match[1])
	}
}
