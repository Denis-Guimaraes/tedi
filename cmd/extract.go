package cmd

import (
	"local/tedi/src/extractor"
	"local/tedi/src/output"

	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract text from files",
	Long:  "Extract text from files.",
	Run: func(cmd *cobra.Command, args []string) {
		e := extractor.New(path, extension, ignore, delimiter)
		extractedTexts := e.ExtractAll()

		o := output.New(outputType)
		o.Write(extractedTexts)
	},
}

func init() {
	extractCmd.Flags().StringVarP(
		&path,
		"path",
		"p",
		"./",
		"folder path to scan",
	)
	extractCmd.Flags().StringSliceVarP(
		&extension,
		"extension",
		"e",
		[]string{".go"},
		"file extension to search",
	)

	extractCmd.Flags().StringSliceVarP(
		&ignore,
		"ignore",
		"i",
		[]string{},
		"glob pattern to ignore folders or files",
	)

	extractCmd.Flags().StringVarP(
		&delimiter,
		"delimiter",
		"d",
		"§",
		"delimiter character",
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
