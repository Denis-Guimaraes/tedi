package output

import "fmt"

func (s *stdout) Write(extractedTexts map[string][]string) {
	for filename, files := range extractedTexts {
		s.writeLines(filename, files)
	}
}

func (s *stdout) writeLines(filename string, texts []string) {
	for _, text := range texts {
		fmt.Printf("%q,%q\n", filename, text)
	}
}
