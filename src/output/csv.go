package output

import "fmt"

func (c *csv) Write(extractedTexts map[string][]string) {
	for filename, files := range extractedTexts {
		c.writeLines(filename, files)
	}
}

func (c *csv) writeLines(filename string, texts []string) {
	for _, text := range texts {
		c.writeLine(filename, text)
	}
}

func (c *csv) writeLine(filename string, text string) {
	fmt.Printf("%q,%q\n", filename, text)
}
