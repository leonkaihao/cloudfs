package cmd

import (
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "watch a directory",
	Long:  "watch a directory to detect any file changes",
}
