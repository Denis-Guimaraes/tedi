package finder

import (
	"fmt"
	"local/tedi/src/logger"
	"os"
	"regexp"
)

func (t *text) Find() []string {
	pattern := fmt.Sprintf("%s(.*?)%s", t.delimiter, t.delimiter)
	regex, err := regexp.Compile(pattern)
	if err != nil {
		logger.Error(fmt.Sprintf("error: %v", err))
		os.Exit(1)
	}
	matches := regex.FindAllStringSubmatch(string(t.content), -1)
	t.addResults(matches)
	return t.results
}

func (t *text) addResults(matches [][]string) {
	for _, match := range matches {
		t.results = append(t.results, match[1])
	}
}
