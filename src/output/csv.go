package output

import (
	ecsv "encoding/csv"
	"fmt"
	"local/tedi/src/logger"
	"os"
)

func (c *csv) Write(extractedTexts map[string][]string) {
	file := c.createFile()
	defer file.Close()
	writer := ecsv.NewWriter(file)
	defer writer.Flush()

	c.writeHeader(writer)
	for filename, texts := range extractedTexts {
		c.writeLines(writer, filename, texts)
	}

	writer.Flush()
	err := writer.Error()
	if err != nil {
		c.error(err)
	}
}

func (c *csv) createFile() *os.File {
	file, err := os.Create("tedi-result.csv")
	if err != nil {
		c.error(err)
	}
	return file
}

func (c *csv) writeHeader(writer *ecsv.Writer) {
	header := []string{"file", "text", "replace"}
	err := writer.Write(header)
	if err != nil {
		c.error(err)
	}
}

func (c *csv) writeLines(writer *ecsv.Writer, filename string, texts []string) {
	for _, text := range texts {
		line := []string{filename, text, ""}
		writer.Write(line)
	}
}

func (c *csv) error(err error) {
	logger.Error(fmt.Sprintf("error: %v", err))
	c.removeFile()
	os.Exit(1)
}

func (c *csv) removeFile() {
	err := os.Remove("tedi-result.csv")
	if err != nil {
		logger.Error(fmt.Sprintf("error: %v", err))
		os.Exit(1)
	}
}
