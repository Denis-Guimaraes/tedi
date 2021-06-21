package cmd

import (
	"local/tedi/src/extractor"
	"local/tedi/src/output"
	"strings"

	"github.com/spf13/cobra"
)

var folder []string
var ignore []string
var pattern string
var outputType string

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract text from files",
	Long:  "Extract text from files.",
	Run: func(cmd *cobra.Command, args []string) {
		e := extractor.New(folder, ignore, strings.Split(pattern, ","))
		extractedTexts := e.ExtractAll()

		o := output.New(outputType)
		o.Write(extractedTexts)
	},
}

func init() {
	extractCmd.Flags().StringSliceVarP(
		&folder,
		"folder",
		"f",
		[]string{"./"},
		"folder path to scan",
	)

	extractCmd.Flags().StringSliceVarP(
		&ignore,
		"ignore",
		"i",
		[]string{},
		"glob pattern to ignore folders or files",
	)

	extractCmd.Flags().StringVarP(
		&pattern,
		"pattern",
		"p",
		"\\{te\\}([^}]*)\\{te\\}",
		"regular expression to extract texts",
	)

	extractCmd.Flags().StringVarP(
		&outputType,
		"output-type",
		"o",
		"stdout",
		"output type, stdout or csv",
	)

	rootCmd.AddCommand(extractCmd)
}
