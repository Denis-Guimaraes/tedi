package cmd

import (
	"local/tedi/src/replacer"

	"github.com/spf13/cobra"
)

var delimiter string
var csv string

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace text in files from previous csv generated",
	Long:  "Replace text in files from previous csv generated",
	Run: func(cmd *cobra.Command, args []string) {
		r := replacer.New(path, ignore, delimiter, csv)
		r.ReplaceAll()
	},
}

func init() {
	replaceCmd.Flags().StringSliceVarP(
		&path,
		"path",
		"",
		[]string{"./"},
		"folder path to scan for find files",
	)

	replaceCmd.Flags().StringSliceVarP(
		&ignore,
		"ignore",
		"",
		[]string{},
		"glob pattern to ignore folders or files",
	)

	replaceCmd.Flags().StringVarP(
		&delimiter,
		"delimiter",
		"",
		"§",
		"delimiter character",
	)

	replaceCmd.Flags().StringVarP(
		&csv,
		"csv",
		"",
		"",
		"csv path",
	)

	replaceCmd.MarkFlagRequired("csv")
	rootCmd.AddCommand(replaceCmd)
}
