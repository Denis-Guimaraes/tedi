package cmd

import (
	"fmt"
	"local/tedi/src/logger"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tedi",
	Short: "Text Editor",
	Long:  "TEdi is a CLI tool writen in go for extract text and replace it.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(fmt.Sprintf("error: %v", err))
		os.Exit(1)
	}
}
