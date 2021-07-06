package cmd

import (
	"local/tedi/src/extractor"
	"local/tedi/src/output"
	"strings"

	"github.com/spf13/cobra"
)

var path []string
var extension []string
var ignore []string
var pattern string
var outputType string

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract text from files",
	Long:  "Extract text from files.",
	Run: func(cmd *cobra.Command, args []string) {
		e := extractor.New(path, extension, ignore, strings.Split(pattern, ","))
		extractedTexts := e.ExtractAll()

		o := output.New(outputType)
		o.Write(extractedTexts)
	},
}

func init() {
	extractCmd.Flags().StringSliceVarP(
		&path,
		"path",
		"",
		[]string{"./"},
		"folder path to scan",
	)
	extractCmd.Flags().StringSliceVarP(
		&extension,
		"extension",
		"",
		[]string{".go"},
		"file extension to search",
	)

	extractCmd.Flags().StringSliceVarP(
		&ignore,
		"ignore",
		"",
		[]string{},
		"glob pattern to ignore folders or files",
	)

	extractCmd.Flags().StringVarP(
		&pattern,
		"pattern",
		"",
		"((?:'|\").*(?:'|\"))",
		"regular expression to extract texts",
	)

	extractCmd.Flags().StringVarP(
		&outputType,
		"output-type",
		"",
		"stdout",
		"output type, stdout or csv",
	)

	rootCmd.AddCommand(extractCmd)
}
