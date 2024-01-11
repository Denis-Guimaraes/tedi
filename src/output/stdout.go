package output

import (
	"fmt"
	"local/tedi/src/logger"
)

func (s *stdout) Write(extractedTexts map[string][]string) {
	for filename, files := range extractedTexts {
		s.writeLines(filename, files)
	}
}

func (s *stdout) writeLines(filename string, texts []string) {
	for _, text := range texts {
		logger.Info(fmt.Sprintf("%q,%q", filename, text))
	}
}
