package cmd

import (
	"fmt"
	"local/tedi/src/extractor"

	"github.com/spf13/cobra"
)

var folder string
var ignore []string
var pattern []string
var output string

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract text from files",
	Long:  "Extract text from files.",
	Run: func(cmd *cobra.Command, args []string) {
		e := extractor.New(folder, ignore, pattern, output)
		extractedTexts := e.ExtractAll()
		fmt.Println(extractedTexts)
	},
}

func init() {
	extractCmd.Flags().StringVarP(
		&folder,
		"folder",
		"f",
		"./",
		"folder path to scan",
	)

	extractCmd.Flags().StringSliceVarP(
		&ignore,
		"ignore",
		"i",
		[]string{},
		"glob pattern to ignore folders or files",
	)

	extractCmd.Flags().StringSliceVarP(
		&pattern,
		"pattern",
		"p",
		[]string{"\\{te\\}([^}]*)\\{te\\}"},
		"regular expression to extract texts",
	)

	extractCmd.Flags().StringVarP(
		&output,
		"output",
		"o",
		"stdout",
		"output type, stdout or csv",
	)

	rootCmd.AddCommand(extractCmd)
}
