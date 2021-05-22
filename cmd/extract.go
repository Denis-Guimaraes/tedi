package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract text from files",
	Long:  "Extract text from files.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Todo")
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}
