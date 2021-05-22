package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Folder string
var Pattern []string
var Output string

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract text from files",
	Long:  "Extract text from files.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Folder)
		for _, p := range Pattern {
			fmt.Println(p)
		}
		fmt.Println(Output)
	},
}

func init() {
	extractCmd.Flags().StringVarP(
		&Folder,
		"folder",
		"f",
		"./",
		"folder path to scan",
	)

	extractCmd.Flags().StringSliceVarP(
		&Pattern,
		"pattern",
		"p",
		[]string{"\\{te\\}([^}]*)\\{te\\}"},
		"regular expression to extract texts",
	)

	extractCmd.Flags().StringVarP(
		&Output,
		"output",
		"o",
		"stdout",
		"output type, stdout or csv",
	)

	rootCmd.AddCommand(extractCmd)
}
