package cmd

import (
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a file/folder to the repo",
	Long:  "Parse a file/folder and record all the file(s) status to the repo",
}
