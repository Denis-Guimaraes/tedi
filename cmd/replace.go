package cmd

import (
	"local/tedi/src/replacer"

	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace text in files from previous csv generated",
	Long:  "Replace text in files from previous csv generated",
	Run: func(cmd *cobra.Command, args []string) {
		r := replacer.New(path, delimiter, csv)
		r.ReplaceAll()
	},
}

func init() {
	replaceCmd.Flags().StringVarP(
		&path,
		"path",
		"p",
		"./",
		"folder path to scan",
	)

	replaceCmd.Flags().StringVarP(
		&delimiter,
		"delimiter",
		"d",
		"§",
		"delimiter character",
	)

	replaceCmd.Flags().StringVarP(
		&csv,
		"csv",
		"c",
		"",
		"csv path",
	)

	replaceCmd.MarkFlagRequired("csv")
	rootCmd.AddCommand(replaceCmd)
}
